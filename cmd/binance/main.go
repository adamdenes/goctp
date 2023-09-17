package main

import (
	"log"

	"github.com/adamdenes/goctp/actors/consumer"
	"github.com/adamdenes/goctp/actors/server"
	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	system := actor.NewActorSystem()

	// Create and start the Server actor
	address := ":8080" // Define the server address and port
	spid, _ := system.Root.SpawnNamed(actor.PropsFromProducer(server.New(address)), "server")
	log.Printf("Starting server actor '%v' with '%v'", spid.Id, spid.Address)

	epid, _ := system.Root.SpawnNamed(actor.PropsFromProducer(consumer.NewBinance()), "binance")
	log.Printf("Starting exchange actor '%v' with '%v'", epid.Id, epid.Address)

	// Block the main thread to keep the actor system running
	// <-make(chan struct{})
	select {}
}
