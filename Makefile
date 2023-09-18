# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -race
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
BINARY_NAME=gocpt

# Protobuf
PROTO_DIR = event

# Scylla/Docker
SCYLLA_IMAGE = scylladb/scylla
SCYLLA_SEED = docker inspect --format='{{ .NetworkSettings.IPAddress }}' $(NODE1_NAME)
NODE1_NAME = scylla-node1
NODE3_NAME = scylla-node2
NODE2_NAME = scylla-node3

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

# Scylla cluster
scylla-start:
	docker run --name $(NODE1_NAME) --hostname $(NODE1_NAME) -d $(SCYLLA_IMAGE) --smp 1 --overprovisioned 1
	docker run --name $(NODE2_NAME) --hostname $(NODE2_NAME) -d $(SCYLLA_IMAGE) \
	 --seeds=`$(SCYLLA_SEED)` --smp 1 --overprovisioned 1
	docker run --name $(NODE3_NAME) --hostname $(NODE3_NAME) -d $(SCYLLA_IMAGE) \
	 --seeds=`$(SCYLLA_SEED)` --smp 1 --overprovisioned 1

scylla-stop:
	@docker stop $(NODE1_NAME) $(NODE2_NAME) $(NODE3_NAME) || true
	@docker rm $(NODE1_NAME) $(NODE2_NAME) $(NODE3_NAME) || true

.PHONY: all build test vet clean run proto scylla-start scylla-stop

