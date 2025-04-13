# AddNet - Microservice Addition Application

AddNet is a distributed microservice application that demonstrates communication between services using gRPC and RabbitMQ message queuing.

## System Architecture

The application consists of two main services:

- **Service A** (Go): A gRPC server that provides an addition operation and forwards the result to a RabbitMQ queue
- **Service B** (TypeScript/Bun): A consumer that reads from the RabbitMQ queue and stores the results, with an HTTP endpoint to retrieve the latest value

### Communication Flow

1. A client makes a gRPC request to Service A with two numbers to add
2. Service A adds the numbers and returns the result to the client
3. Service A simultaneously sends the result to a RabbitMQ queue
4. Service B consumes the message from the queue and persists the result
5. The latest result can be retrieved via Service B's HTTP endpoint

## Services

### Service A (Go)

A Go-based gRPC service that:

- Exposes an `Add` method through gRPC
- Publishes results to the `addition_queue` in RabbitMQ
- Exposes Prometheus metrics on port 2112

#### To run Service A:

```bash
cd service_a
go run main.go
```

### Service B (TypeScript/Bun)

A TypeScript service built with Bun that:

- Consumes messages from the `addition_queue` in RabbitMQ
- Persists the running total to a file
- Provides a REST endpoint at `http://localhost:3000/latest` to retrieve the current value

#### To run Service B:

```bash
cd service_b
bun install
bun run index.ts
```

## Prerequisites

- Go 1.24+
- Bun 1.2.9+
- RabbitMQ running on localhost:5672
- gRPC tools (for development)

## Configuration

Both services are configured to connect to RabbitMQ at `amqp://localhost:5672` by default.

## API

### gRPC (Service A)

The gRPC service exposes the following RPC:

```proto
service AdditionService {
    rpc Add(AddRequest) returns (AddResponse) {}
}
```

### REST (Service B)

- `GET /latest` - Returns the current accumulated value

## Development

To regenerate the gRPC code after modifying the protobuf definition:

```bash
cd service_a
protoc --go_out=. --go-grpc_out=. addition.proto
```
