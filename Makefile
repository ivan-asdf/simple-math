build:
	docker compose up --build

run:
	docker compose up

build_local:
	go build -o build/api ./cmd/api/main.go
	go build -o build/cli ./cmd/cli/main.go
	cp .env build/.env

.PHONY: build run build_local
