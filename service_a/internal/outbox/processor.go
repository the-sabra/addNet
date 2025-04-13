package outbox

import (
	"addition_service/internal/ports/rabbitmq"
	"context"
	"fmt"
	"log"
	"time"
)

type Processor struct {
	Repository Repository
	publisher  *rabbitmq.Publisher
	batchSize  int
	interval   time.Duration
}

func NewProcessor(repository Repository, publisher *rabbitmq.Publisher, batchSize int, interval time.Duration) *Processor {
	return &Processor{
		Repository: repository,
		publisher:  publisher,
		batchSize:  batchSize,
		interval:   interval,
	}
}

func (p *Processor) Start(ctx context.Context) {
	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()

	log.Println("Starting outbox processor")

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down outbox processor")
			return
		case <-ticker.C:
			if err := p.processMessages(ctx); err != nil {
				log.Printf("Error processing outbox messages: %v", err)
			}
		}
	}
}

func (p *Processor) processMessages(ctx context.Context) error {
	log.Println("Fetching pending messages from outbox")
	outboxes, err := p.Repository.GetPendingMessages(ctx, p.batchSize)
	if err != nil {
		log.Println("Error fetching pending messages:", err)
		return err
	}

	for _, outbox := range outboxes {
		err = p.publisher.PublishMessage("addition_queue", []byte(fmt.Sprintf("%d", outbox.Value)))
		if err != nil {
			log.Printf("Error publishing message to RabbitMQ: %v", err)
			return err
		}

		err = p.Repository.MarkAsProcessed(ctx, outbox.ID)
		if err != nil {
			log.Printf("Error marking message as processed: %v", err)
			return err
		}
	}

	return nil
}
