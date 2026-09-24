.PHONY: help build run run-dev lint migrate-up migrate-down podman-up podman-down

help:
	@echo Available commands:
	@echo  build     		- Build the project
	@echo  run       		- Run the project
	@echo  run-dev   		- Run the project in development mode
	@echo  lint      		- Lint the project
	@echo  migrate-up 		- Run database migrations
	@echo  migrate-down 	- Rollback database migrations
	@echo  podman-up 		- Start Podman containers
	@echo  podman-down  	- Stop Podman containers

build:
	@echo "Building the project..."
	go build -o ./bin/GoCart/go-cart.exe ./cmd/api/main.go

run:
	@echo "Running the project..."
	@./bin/GoCart/go-cart.exe

run-dev:
	@echo "Running the project in development mode..."
	go tool air

lint:
	@echo "Linting the project..."
	golangci-lint run

migrate-up:
	migrate -path db/migrations -database "postgresql://gcdba:gcpwd@localhost:5432/gcdb?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgresql://gcdba:gcpwd@localhost:5432/gcdb?sslmode=disable" down

podman-up:
	@echo "Starting Podman containers..."
	podman compose -f container/podman-compose.yaml up -d

podman-down:
	@echo "Stopping Podman containers..."
	podman compose -f container/podman-compose.yaml down
