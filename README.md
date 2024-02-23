# Project Structure

This Go project follows a clean architecture approach, with a clear separation of concerns and dependencies. Here's a brief overview of the directory structure:

- `api/`: Contains the OpenAPI specification for the API.
- `cmd/`: Contains the application's entry point (`main.go`). This is where the application is bootstrapped and run.
- `config/`: Contains the application's configuration files. The `config.go` file loads environment variables.
- `internal/`: Contains the business logic of the application. It's further divided into:
  - `api/`: Contains the HTTP handlers, middlewares, and routes.
  - `domain/`: Contains the domain models and services.
  - `errors/`: Contains custom error types for the application.
  - `spi/`: Contains the service provider interfaces.
  - `utils/`: Contains utility functions and helpers.
- `pkg/`: Contains packages that can be used by other services.
- `scripts/`: Contains scripts for tasks like generating the OpenAPI specification.
- `migrations/`: Contains SQL migration scripts.
- `tmp/`: Contains temporary files generated during development.

## Key Files

- `main.go`: The entry point of the application. It sets up the HTTP server, database connection, and routes.
- `config.go`: Loads the application's configuration from environment variables.
- `openapi-spec-generator.go`: A script that generates the OpenAPI specification for the API.
- `Dockerfile`: Defines how to build a Docker image for the application.
- `Makefile`: Contains commands for building, running, and testing the application.
- `docker-compose.yml`: Used to define and run multi-container Docker applications.
- `go.mod` and `go.sum`: Used to manage the application's dependencies.

## Running the Application

To run the application, use the `make up` command. This will build and run the application in a Docker container. Use `make down` to stop the application.

## Testing

Unit tests can be run with `make unit-test`. Integration tests can be run with `make integration-test`.

## Building

To build the application, use the `make build` command. This will create a Docker image for the application.

## Migrations

Database migrations are handled with the `migrate` command. Use `make migrate-down` to run down migrations.

For more details, please refer to the individual files and code comments.