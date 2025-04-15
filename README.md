Here's a comprehensive README.md file for your project, giving credit to Eng Ahmed El Taweel and including all the key information:

# AddNet - Distributed Microservice Application

[![Go](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/)
[![Bun](https://img.shields.io/badge/Bun-1.2.9+-yellow.svg)](https://bun.sh/)
[![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.9+-orange.svg)](https://www.rabbitmq.com/)
[![Prometheus](https://img.shields.io/badge/Prometheus-2.33+-red.svg)](https://prometheus.io/)
[![Grafana](https://img.shields.io/badge/Grafana-8.5+-purple.svg)](https://grafana.com/)

AddNet is a distributed microservice application demonstrating communication between services using gRPC and RabbitMQ message queuing, inspired by [Eng Ahmed El Taweel video](https://youtu.be/Ur6b1NWGbYE?si=54EkJ_Avy5QT-YGF).

## System Architecture

![Architecture Diagram](https://i.imgur.com/gztFnNk.png)

The application consists of two main services:

- **Service A** (Go): A gRPC server that provides an addition operation and forwards results to RabbitMQ
- **Service B** (TypeScript/Bun): A consumer that reads messages from RabbitMQ, adds the new value to an existing total, saves the updated value to a file, and exposes an HTTP endpoint to retrieve the latest result

### Communication Flow

1. Client makes gRPC request to Service A with two numbers
2. Service A adds numbers and returns result
3. Service A sends result to RabbitMQ queue
4. Service B consumes messages and persists results
5. Latest result available via Service B's HTTP endpoint

## Features

- **Reliable Messaging**: Outbox pattern ensures message delivery
- **Monitoring**: Prometheus metrics and Grafana dashboard
- **Load Testing**: k6 test scripts included
- **Containerized**: Docker Compose for easy deployment
- **Multi-language**: Go and TypeScript services

## Monitoring

![Monitoring Diagram](https://i.imgur.com/MOaO8Pi.jpeg)

## Prerequisites

- Docker and Docker Compose
- Go 1.24+
- Bun 1.2.9+ (for Service B)
- gRPC tools (for development)

## Getting Started

### Running with Docker Compose

```bash
docker-compose up --build
```

This will start:

- Service A (gRPC server on port 50051)
- Service B (HTTP server on port 2113)
- RabbitMQ (management UI on port 15672)
- PostgreSQL database
- Prometheus (on port 9090)
- Grafana (on port 3000)

### Accessing Services

- **Service A gRPC**: `localhost:50051`
- **Service B HTTP**: `http://localhost:2113/latest`
- **RabbitMQ Management**: `http://localhost:15672` (guest/guest)
- **Prometheus**: `http://localhost:9090`
- **Grafana**: `http://localhost:3000` (admin/grafana)

## Development

### Service A (Go)

```bash
cd service_a
go run cmd/server/main.go
```

### Service B (TypeScript)

```bash
cd service_b
bun install
bun run index.ts
```

### Regenerating gRPC Code

After modifying `addition.proto`:

```bash
cd service_a
protoc --go_out=. --go-grpc_out=. internal/proto/addition.proto
```

## Testing

### Load Testing with k6

**gRPC Test (Service A):**

```bash
k6 run service_a/k6_grpc_test.js
```

**REST Test (Service B):**

```bash
k6 run service_b/k6_rest_test.js
```

## Monitoring

The project includes preconfigured Grafana dashboards showing:

- Request latency
- Throughput
- Error rates
- Queue delivery metrics

## Project Structure

```
the-sabra-addnet/
├── service_a/          # Go gRPC service
│   ├── cmd/           # Application entry points
│   ├── internal/      # Internal application code
│   └── proto/         # Protocol Buffer definitions
├── service_b/         # TypeScript consumer service
├── prometheus/        # Prometheus configuration
└── grafana/           # Grafana dashboards and datasources
```

## Acknowledgments

This project was inspired by Eng Ahmed El Taweel excellent tutorial on building distributed systems. Check out his [original video](https://youtu.be/Ur6b1NWGbYE?si=54EkJ_Avy5QT-YGF) for more insights into microservice architecture.
