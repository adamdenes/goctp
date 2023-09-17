package main

import (
	"log"

	"github.com/adamdenes/goctp/actors/consumer"
	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	system := actor.NewActorSystem()
	pid, _ := system.Root.SpawnNamed(actor.PropsFromProducer(consumer.NewBinance()), "binance")
	log.Printf("Starting actor %v with %v", pid.Id, pid.Address)
	select {}
}
