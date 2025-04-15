package outbox

import (
	"addition_service/internal/ports/rabbitmq"
	"context"
	"fmt"
	"log"
	"sync"
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
	if len(outboxes) == 0 {
		return nil
	}

	log.Printf("Processing %d outbox messages concurrently", len(outboxes))

	var wg sync.WaitGroup
	errChan := make(chan error, len(outboxes))

	for _, outbox := range outboxes {
		wg.Add(1)
		go func(ob Outbox) {
			defer wg.Done()

			if err := p.publisher.PublishMessage("addition_queue", ob.Value); err != nil {
				errChan <- fmt.Errorf("publish message ID %d: %w", ob.ID, err)
				return
			}

			if err := p.Repository.MarkAsProcessed(ctx, ob.ID); err != nil {
				errChan <- fmt.Errorf("mark message ID %d as processed: %w", ob.ID, err)
				return
			}
		}(outbox)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var allErrors []error
	for err := range errChan {
		log.Printf("Error processing outbox message: %v", err)
		allErrors = append(allErrors, err)
	}

	if len(allErrors) > 0 {
		return fmt.Errorf("encountered errors: %v", allErrors)
	}

	return nil
}
