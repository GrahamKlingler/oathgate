COMPOSE := docker compose -f deploy/compose/docker-compose.yml

.PHONY: build test lint run up down

build:
	go build -o bin/ ./cmd/...

test:
	go test -race ./...

lint:
	golangci-lint run

run:
	go run ./cmd/gateway

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down
