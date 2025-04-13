package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"addition_service/internal/db"
	"addition_service/internal/outbox"
	"addition_service/internal/ports/grpc"
	"addition_service/internal/ports/rabbitmq"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to RabbitMQ
	conn, ch, err := rabbitmq.NewRabbitMQClient()
	defer cancel()

	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
		return
	}

	err = rabbitmq.DeclareQueue(ch, "addition_queue")
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
		return
	}

	defer ch.Close()
	defer conn.Close()
	// start the db
	db := db.InitDB()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	defer db.Close()

	outboxRepo := outbox.NewSQLRepository(db)
	publisher := rabbitmq.NewPublisher(ch)

	outboxProcessor := outbox.NewProcessor(
		outboxRepo,
		publisher,
		10,            // process 10 messages at a time
		5*time.Second, // process every 5 seconds
	)

	go outboxProcessor.Start(ctx)

	// Call the GRPC server and keep the main function running
	server := grpc.NewGRPCServer(outboxRepo)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the HTTP server for Prometheus metrics in a separate goroutine
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting metrics server on port :2112...")
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatalf("Failed to start metrics server: %v", err)
		}
	}()

	log.Println("Starting gRPC server on port :50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
