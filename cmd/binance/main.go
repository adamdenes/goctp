package main

import (
	"fmt"

	"github.com/adamdenes/goctp/actors/consumer"
	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	system := actor.NewActorSystem()
	pid, _ := system.Root.SpawnNamed(actor.PropsFromProducer(consumer.NewBinance()), "binance")
	fmt.Println(pid)
	select {}
}
