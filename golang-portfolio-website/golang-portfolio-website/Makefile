.PHONY: run build test clean migrate deps

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

test:
	go test ./... -v

clean:
	rm -rf bin/

migrate:
	psql -d portfolio -f migrations/init.sql

deps:
	go mod download
	go mod tidy
