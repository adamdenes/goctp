# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -race
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
BINARY_NAME=gocpt

PROTO_DIR = event

SCYLLA_IMAGE = scylladb/scylla
NODE1_NAME = scylla-node1

all: build

build:
	$(GOBUILD) -o ./bin/app/$(BINARY_NAME) -v cmd/binance/main.go

test:
	$(GOTEST) ./...

vet:
	$(GOVET) ./...

clean:
	$(GOCLEAN)
	rm -f ./bin/app/$(BINARY_NAME)

run:
	./bin/app/$(BINARY_NAME)

# Protobuff compilation
proto:
	protoc --proto_path=$(PROTO_DIR) --go_out=$(PROTO_DIR) event.proto;

# Docker
docker-start:
	docker run --name $(NODE1_NAME) --hostname $(NODE1_NAME) -d $(SCYLLA_IMAGE) --smp 1 --overprovisioned 1

docker-stop:
	@docker stop $(NODE1_NAME) || true
	@docker rm $(NODE1_NAME) || true

.PHONY: all build test vet clean run proto docker-start docker-stop

