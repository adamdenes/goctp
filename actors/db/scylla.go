package db

import (
	"fmt"
	"log"

	"github.com/adamdenes/goctp/event"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/gocql/gocql"
)

type ScyllaDB struct {
	session  *gocql.Session
	hosts    []string
	keyspace string
}

func NewScyllaDB(h []string, k string) actor.Producer {
	return func() actor.Actor {
		return &ScyllaDB{
			hosts:    h,
			keyspace: k,
		}
	}
}

func (a *ScyllaDB) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		// Perform initialization tasks here, e.g., creating tables,
		// preparing statements
		a.start()
	case *event.Kline:
		// log.Println("[RECEIVED]:", msg)
		// Create a table for a specific symbol/actor
		symbol := msg.GetSymbol()
		if err := a.createTableForSymbol(symbol); err != nil {
			log.Fatalf("Failed to create table: %v\n", err)
		}
		// Insert data into the corresponding table
		a.insertKlineData(symbol, msg)

	case *event.HistoricalKline:
		// TODO: this event will provide Candlestick information from the REST API
		// and this data will be used for backtesting as opposed to the real-time
		// websocket data (*event.Kline).
		// HistoricalKline needs to be stored in a separate table. Once saved,
		// Plot(ter) actor will fetch it and plot it on a webinterface
	case *actor.Stopped:
		a.session.Close()
	}
}

func (a *ScyllaDB) start() {
	log.Println("ScyllaDB actor started")
	session, err := a.connect()
	if err != nil {
		log.Fatalf("Failed to create ScyllaDB session: %v\n", err)
	}
	a.session = session

	// Attempt to create the keyspace with `IF NOT EXISTS``
	// otherwise connection FAILS!!!
	if err := a.createKeyspace(a.session); err != nil {
		log.Fatalf("Failed to create keyspace: %v\n", err)
	}
}

func (a *ScyllaDB) connect() (*gocql.Session, error) {
	log.Println("Connecting to database...")
	// Initialize database config
	cluster := gocql.NewCluster(a.hosts...)
	cluster.Consistency = gocql.Quorum
	cluster.Port = 9042
	cluster.ProtoVersion = 0

	// Create new session
	return cluster.CreateSession()
}

func (a *ScyllaDB) createKeyspace(session *gocql.Session) error {
	createKeyspaceQuery := fmt.Sprintf(`
		CREATE KEYSPACE IF NOT EXISTS %s
		WITH replication = {
			'class': 'SimpleStrategy',
			'replication_factor': 1
		};
	`, a.keyspace)

	if err := session.Query(createKeyspaceQuery).Exec(); err != nil {
		return fmt.Errorf("error creating keyspace %s: %v", a.keyspace, err)
	}
	return nil
}

func (a *ScyllaDB) createTableForSymbol(symbol string) error {
	// Define your table creation query here, reflecting Kline data structure
	createTableQuery := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s.%s (
			eventType text,
			eventTime bigint,
			symbol text,
			startTime bigint,
			closeTime bigint,
			symbolInner text,
			interval text,
			firstTradeID bigint,
			lastTradeID bigint,
			openPrice double,
			closePrice double,
			highPrice double,
			lowPrice double,
			baseAssetVolume bigint,
			numberOfTrades bigint,
			isKlineClosed boolean,
			quoteAssetVolume double,
			takerBuyBaseAssetVolume bigint,
			takerBuyQuoteAssetVolume double,
			ignore bigint,
			PRIMARY KEY (eventTime, symbol)
		);
	`, a.keyspace, symbol)

	if err := a.session.Query(createTableQuery).Exec(); err != nil {
		return fmt.Errorf("error creating table for symbol %s: %v", symbol, err)
	}
	return nil
}

func (a *ScyllaDB) insertKlineData(symbol string, klineData *event.Kline) error {
	// INSERT statement with placeholders
	insertStmt := fmt.Sprintf(`
		INSERT INTO %s.%s (
			eventType, eventTime, symbol,
			startTime, closeTime, symbolInner,
			interval, firstTradeID, lastTradeID,
			openPrice, closePrice, highPrice,
			lowPrice, baseAssetVolume, numberOfTrades,
			isKlineClosed, quoteAssetVolume,
			takerBuyBaseAssetVolume, takerBuyQuoteAssetVolume,
			ignore
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`, a.keyspace, symbol)

	// Execute the INSERT statement with the marshaled data
	err := a.session.Query(insertStmt).Bind(
		klineData.EventType, klineData.EventTime, klineData.Symbol,
		klineData.StartTime, klineData.CloseTime, klineData.SymbolInner,
		klineData.Interval, klineData.FirstTradeID, klineData.LastTradeID,
		klineData.OpenPrice, klineData.ClosePrice, klineData.HighPrice,
		klineData.LowPrice, klineData.BaseAssetVolume, klineData.NumberOfTrades,
		klineData.IsKlineClosed, klineData.QuoteAssetVolume,
		klineData.TakerBuyBaseAssetVolume, klineData.TakerBuyQuoteAssetVolume,
		klineData.Ignore,
	).Exec()

	if err != nil {
		return fmt.Errorf("error executing INSERT statement: %v", err)
	}
	return nil
}
