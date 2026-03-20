run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

test:
	go test ./...

lint:
	golangci-lint run

migrate-up:
	migrate -path migrations -database $$DATABASE_URL up

migrate-down:
	migrate -path migrations -database $$DATABASE_URL down