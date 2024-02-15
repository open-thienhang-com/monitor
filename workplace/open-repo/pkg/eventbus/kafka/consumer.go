package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"mono.thienhang.com/pkg/eventbus/base"
)

type KafkaConsumer struct {
	reader *kafka.Reader
	// pool              *ants.Pool
	releaseTimeoutSec time.Duration
}

func NewKafkaEventConsumer(cfg *ConsumerConfig) base.EventConsumer {
	consumer := &KafkaConsumer{}
	consumer.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cfg.Brokers,
		GroupID:  cfg.GroupID,
		Topic:    cfg.Topic,
		MaxWait:  3 * time.Second,
		MinBytes: 1, // read 1 byte at a time
		MaxBytes: 10e6,
	})
	return consumer
}

func (c *KafkaConsumer) Consume(ctx context.Context, handler func(*Event) error) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v\n", err)
				continue
			}

			event := &Event{
				ID:   string(msg.Key),
				Name: string(msg.Value),
				// Extract more fields as needed
			}

			if err := handler(event); err != nil {
				log.Printf("Error handling message: %v\n", err)
			}
		}
	}
}

func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}

func (c *KafkaConsumer) SafeClose() error {
	return c.reader.Close()
}
