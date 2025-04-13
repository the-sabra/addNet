package main

import (
	"log"
	"net"

	"net/http"

	pb "addition_service/server/addition"

	"addition_service/server"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	log.Println("Connected to RabbitMQ")

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
		return
	}

	defer ch.Close()

	// Declare a queue
	_, err = ch.QueueDeclare(
		"addition_queue", // name
		false,            // delete when unused
		false,            // exclusive
		false,
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare a queue: %v", err)
		return
	}

	log.Println("Queue declared")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	reflection.Register(grpcServer)

	pb.RegisterAdditionServiceServer(grpcServer, server.NewAdditionService(ch))

	// Start the HTTP server for Prometheus metrics in a separate goroutine
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting metrics server on port :2112...")
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatalf("Failed to start metrics server: %v", err)
		}
	}()

	// Start the gRPC server (this is a blocking call)
	log.Println("Starting gRPC server on port :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
