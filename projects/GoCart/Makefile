.PHONY: help build run run-build run-dev fmt lint migrate-up migrate-down docker-up docker-down

help:
	@echo Available commands:
	@echo  build     		- Build the project
	@echo  run       		- Run the project
	@echo  run-build 		- Run the built project
	@echo  run-dev   		- Run the project in development mode
	@echo  fmt       		- Format the code
	@echo  lint      		- Lint the project
	@echo  migrate-up 		- Run database migrations
	@echo  migrate-down 	- Rollback database migrations
	@echo  docker-up 		- Start docker containers
	@echo  docker-down  	- Stop docker containers

build:
	@echo "Building the project..."
	go build -o ./bin/GoCart/go-cart.exe ./cmd/api/main.go

run:
	@echo "Running the project..."
	go run ./cmd/api/main.go

run-build:
	@echo "Running the project..."
	@./bin/GoCart/go-cart.exe

run-dev:
	@echo "Running the project in development mode..."
	go tool air

fmt:
	@echo "Formatting the code..."
	go fmt ./...

lint:
	@echo "Linting the project..."
	golangci-lint run

migrate-up:
	migrate -path db/migrations -database "postgresql://gcdba:gcpwd@localhost:5433/gcdb?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgresql://gcdba:gcpwd@localhost:5433/gcdb?sslmode=disable" down

docker-up:
	@echo "Starting docker containers..."
	docker compose -f docker/docker-compose.yaml up -d

docker-down:
	@echo "Stopping docker containers..."
	docker compose -f docker/docker-compose.yaml down
