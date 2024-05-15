build:
	docker compose up --build

build_local:
	go build -o build/api ./cmd/api/main.go
	go build -o build/cli ./cmd/cli/main.go
	cp .env build/.env
