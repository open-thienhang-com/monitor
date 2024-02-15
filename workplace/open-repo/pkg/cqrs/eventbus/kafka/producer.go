package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"mono.thienhang.com/pkg/eventbus/base"
)

type KafkaProducer struct {
	writer *kafka.Writer
	// pool              *ants.Pool
	releaseTimeoutSec time.Duration
}

// NewProducer create new producer with input configuration
func NewProducer(cfg *ProducerConfig) (base.EventBus, error) {
	// if cfg.Mock {
	// 	return &Mock{}, nil
	// }

	producer := &KafkaProducer{
		// cfg: cfg,
	}

	// var err error
	// producer.pool, err = ants.NewPool(cfg.PoolSize, ants.WithLogger(log))
	// if err != nil {
	// 	return nil, err
	// }

	producer.writer = &kafka.Writer{
		Addr:         kafka.TCP(cfg.Brokers...),
		Async:        true,
		BatchTimeout: 10 * time.Millisecond,
		Balancer:     &kafka.LeastBytes{},
	}

	if cfg.ReleaseTimeoutSec > 0 {
		producer.releaseTimeoutSec = time.Duration(cfg.ReleaseTimeoutSec) * time.Second
	} else {
		producer.releaseTimeoutSec = time.Minute * 10
	}

	return producer, nil
}

// Stop the consumer and close deadletter
func (p *KafkaProducer) Stop() {
	// Need to wait all message to be finish before stop othewise we will losing data
	// p.pool.ReleaseTimeout(p.releaseTimeoutSec)
}

// Publish write input message to destination topic on kafka
func (p *KafkaProducer) Publish(topic string, msg interface{}) error {
	// p.pool.Submit(func() {
	// 	// verify and mashal message if need
	// 	encodedMsg, ok := msg.([]byte)
	// 	if !ok {
	// 		encodedMsgByte, err := json.Fast.Marshal(msg)
	// 		if err != nil {
	// 			logger.Error("Could not encode message", logger.Field("message", msg), logger.Field("error", err))
	// 			return
	// 		}
	// 		encodedMsg = encodedMsgByte
	// 	}

	// 	// async write message to kafka
	// 	p.writer.WriteMessages(context.Background(), kafka.Message{
	// 		Topic: topic,
	// 		Value: encodedMsg,
	// 	})
	// })
	return nil
}

func (p *KafkaProducer) PublishRaw(topic string, msg []byte) error {
	// async write message to kafka
	err := p.writer.WriteMessages(context.TODO(), kafka.Message{
		Topic: topic,
		Value: msg,
	})

	return err
}
