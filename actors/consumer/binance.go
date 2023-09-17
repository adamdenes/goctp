package consumer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/adamdenes/goctp/actors/db"
	"github.com/adamdenes/goctp/event"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/valyala/fastjson"
	"nhooyr.io/websocket"
)

const wsEndpoint = "wss://stream.binance.com/stream?streams=btcusdt@markPrice"
const apiKey = "6r3PLGC5RcRnHIlMkAej55otVT9YHPPkXKCB4z2dUIDx698MUVj1IvOcQPBnEFns"
const apiSecret = "xVdliUOTP48qj0gQj96MKJ6F8Vmf4urj2vVEXSwlODINMxDXA8tXXBzf307Qt3q2"

// figure out a better way for this
var hosts = []string{"172.17.0.2"}

// temporary selection, will be configured from webinterface
var symbols = []string{"btcusdt", "ethusdt"}

// make it variadic
func createWsEndpoint() string {
	results := make([]string, 0, 3)
	for _, symbol := range symbols {
		results = append(results, fmt.Sprintf("%s@aggTrade", symbol))
		results = append(results, fmt.Sprintf("%s@kline_1s", symbol))
		results = append(results, fmt.Sprintf("%s@depth", symbol))
	}
	return fmt.Sprintf("%s%s", wsEndpoint, strings.Join(results, "/"))
}

type Binance struct {
	ws    *websocket.Conn
	ctx   context.Context
	actx  actor.Context
	dbPid map[string]*actor.PID
}

func NewBinance() actor.Producer {
	return func() actor.Actor {
		return &Binance{
			dbPid: make(map[string]*actor.PID),
		}
	}
}

func (b *Binance) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		_ = msg
		b.ctx = context.Background()
		b.actx = c
		b.start()
	case *actor.Stopped:
		b.ws.Close(websocket.StatusAbnormalClosure, "Closed by client")
	}
}

func (b *Binance) msgLoop() {
	for {
		// TODO: might not need to increase size?
		b.ws.SetReadLimit(65536)
		_, msg, err := b.ws.Read(b.ctx)
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			log.Println(err)
			continue
		}
		parser := fastjson.Parser{}
		v, err := parser.ParseBytes(msg)
		if err != nil {
			log.Println(err)
			continue
		}
		stream := v.GetStringBytes("stream")
		symbol, kind := splitStream(string(stream))
		data := v.Get("data")

		// fmt.Println(data.GetObject("k").Get("o"))
		// fmt.Println(symbol, "->", kind)

		switch kind {
		case "aggTrade":
		case "depth":
		case "kline_1s":
			if pid, ok := b.dbPid[symbol]; ok {
				kline := composeKline(data)
				b.actx.Send(pid, kline)
			}
		}
	}
}

func (b *Binance) start() {
	log.Println("Binance/WebSocket actor started")
	header := make(http.Header)
	header.Add("APCA-API-KEY-ID", apiKey)
	header.Add("APCA-API-SECRET-KEY", apiSecret)

	endpoints := createWsEndpoint()

	conn, _, err := websocket.Dial(b.ctx, endpoints, &websocket.DialOptions{
		HTTPHeader: header,
	})
	if err != nil {
		log.Fatal("dial:", err)
	}
	b.ws = conn

	for _, s := range symbols {
		pid, _ := b.actx.SpawnNamed(actor.PropsFromProducer(db.NewScyllaDB(hosts, "binance")), s)
		b.dbPid[s] = pid
		log.Printf("Spwaning scylla child actor: %v\n", b.dbPid[s].Id)
	}

	go b.msgLoop()
}

func splitStream(stream string) (string, string) {
	parts := strings.Split(stream, "@")
	return parts[0], parts[1]
}

func composeKline(data *fastjson.Value) *event.Kline {
	d := data.GetObject("k")
	return &event.Kline{
		EventType:                string(data.GetStringBytes("e")),
		EventTime:                data.GetInt64("E"),
		Symbol:                   string(data.GetStringBytes("s")),
		StartTime:                d.Get("t").GetInt64(),
		CloseTime:                d.Get("T").GetInt64(),
		SymbolInner:              d.Get("s").String(),
		Interval:                 d.Get("i").String(),
		FirstTradeID:             d.Get("f").GetInt64(),
		LastTradeID:              d.Get("L").GetInt64(),
		OpenPrice:                d.Get("o").GetFloat64(),
		ClosePrice:               d.Get("c").GetFloat64(),
		HighPrice:                d.Get("h").GetFloat64(),
		LowPrice:                 d.Get("l").GetFloat64(),
		BaseAssetVolume:          d.Get("v").GetInt64(),
		NumberOfTrades:           d.Get("n").GetInt64(),
		IsKlineClosed:            d.Get("x").GetBool(),
		QuoteAssetVolume:         d.Get("q").GetFloat64(),
		TakerBuyBaseAssetVolume:  d.Get("V").GetInt64(),
		TakerBuyQuoteAssetVolume: d.Get("Q").GetFloat64(),
		Ignore:                   d.Get("B").GetInt64(),
	}
}
