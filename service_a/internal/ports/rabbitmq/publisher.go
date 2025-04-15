package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Message struct {
	Value  int32  `json:"value"`
	SentAt string `json:"sent_at"`
}

// Publisher is a struct that holds the RabbitMQ connection and channel.
type Publisher struct {
	ch *amqp.Channel
}

func NewPublisher(ch *amqp.Channel) *Publisher {
	return &Publisher{ch: ch}
}

func (pb *Publisher) PublishMessage(queueName string, value int32) error {
	message := &Message{
		Value:  value,
		SentAt: time.Now().Format(time.RFC3339),
	}

	body, err := json.Marshal(message)

	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = pb.ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Published message to queue: %s", queueName)
	return nil
}
