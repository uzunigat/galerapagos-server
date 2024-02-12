FROM golang:1.20 AS build


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o bin/service ./cmd

FROM debian:buster-slim AS app

RUN apt-get update && \
    apt-get install -y ca-certificates

WORKDIR /app
COPY --from=build /app/bin/service .

COPY migrations migrations

EXPOSE 3000
CMD ["./service"]