package server

import (
	"addition_service/server/addition"
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AdditionService struct {
	addition.UnimplementedAdditionServiceServer
	ch *amqp.Channel
}

func NewAdditionService(ch *amqp.Channel) *AdditionService {
	return &AdditionService{
		ch: ch,
	}
}

func (s *AdditionService) Add(ctx context.Context, req *addition.AddRequest) (*addition.AddResponse, error) {
	result := req.GetA() + req.GetB()
	s.addToQueue(result)
	return &addition.AddResponse{Result: result}, nil
}

func (s *AdditionService) addToQueue(result int32) {
	// Implement the logic to add the result to RabbitMQ queue here
	// For example, you can use the channel to publish a message to the queue
	err := s.ch.Publish(
		"",
		"addition_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%d", result)),
		},
	)
	if err != nil {
		log.Fatalf("failed to publish a message: %v", err)
		return
	}
	log.Printf("Published message: %d", result)
}
