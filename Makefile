.PHONY: audit build clean lint migrate test-local test-unit test-integration seed assume-role codeartifact-token
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build:
	@echo "---> 🏗️ Building go service..."
	@docker-compose build app

up: build
	@echo "---> Running service container..."
	@docker-compose up app
	@echo "Exited."

up-dev:
	@echo "---> Starting dependencies..."
	@docker-compose up app-dev || (echo "Failed to start app-dev container"; exit 1)
	@echo "---> starting service and watching for changes"
	@echo "Exited."

up-local-pulsar: build
	@echo "---> Running service container with local pulsar cluster..."
	@docker-compose up app pulsar
	@echo "Exited."

down:
	@echo "---> Stopping service..."
	@docker-compose down 

clean:
	@echo "---> 🧹 Cleaning up some things..."
	@echo "---> 🐳 Stopping and removing docker artifacts..."
	@docker-compose down -v --remove-orphans --rmi local
	@echo "---> 🗑  Deleting miscellaneous artifacts..."
	@rm -rf *.log coverage/ dist/
	@echo "---> ✅ Done"

generate-api-spec:
	@echo "Generating open api spec..."
	@go run ./scripts/openapi-spec-generator/openapi-spec-generator.go

vet:
	@echo "🩺  Vetting code via docker"
	@docker-compose run --no-deps app go vet ./...

migrate-down: 
	@echo "---> 🦆 Running migration scripts..."
	@docker compose up db -d
	@echo "dbname postgres://${RDS_USERNAME}:${RDS_PASSWORD}@localhost:${RDS_PORT}/${RDS_DBNAME}?sslmode=disable" 
	@migrate -source file://./migrations -database "postgres://${RDS_USERNAME}:${RDS_PASSWORD}@localhost:${RDS_PORT}/${RDS_DBNAME}?sslmode=disable" down
	@docker compose down


unit-test-local: 
	@echo "---> 🦄 Running unit tests locally..."	
	@go test ./... -v --tags=unit

unit-test:
	@echo "---> 🦄 Running unit tests..."
	@docker-compose up --build --exit-code-from unit-test unit-test

integration-test:
	@echo "---> 🌎 Running integration tests..."
	@docker-compose up --build --exit-code-from integration-test integration-test

tidy-deps:
	@echo "---> 🧹 Tidying up module imports..."
	@go mod tidy