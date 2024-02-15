package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockProducerConfig is a mock configuration for the KafkaProducer.
var MockProducerConfig = &ProducerConfig{
	Brokers:           []string{"localhost:9092"},
	Topic:             "test-topic",
	Replication:       1,
	Partition:         0,
	BufferSize:        100,
	PoolSize:          5,
	ReleaseTimeoutSec: 10,
}

func TestNewProducer(t *testing.T) {
	// Test NewProducer function with mock configuration
	producer, err := NewProducer(MockProducerConfig)
	assert.NoError(t, err)
	assert.NotNil(t, producer)
	defer producer.Close()
}

func TestKafkaProducer_Publish(t *testing.T) {
	// Create a KafkaProducer with mock configuration
	producer, err := NewProducer(MockProducerConfig)
	assert.NoError(t, err)
	assert.NotNil(t, producer)
	defer producer.Close()

	// Mock data
	topic := "test-topic"
	message := "test-message"

	// Test the Publish method
	err = producer.Publish(topic, message)
	assert.NoError(t, err)
}

func TestKafkaProducer_PublishRaw(t *testing.T) {
	// Create a KafkaProducer with mock configuration
	producer, err := NewProducer(MockProducerConfig)
	assert.NoError(t, err)
	assert.NotNil(t, producer)
	defer producer.Close()

	// Mock data
	topic := "test-topic"
	message := []byte("test-message")

	// Test the PublishRaw method
	err = producer.PublishRaw(topic, message)
	assert.NoError(t, err)
}

func TestKafkaProducer_Close(t *testing.T) {
	// Create a KafkaProducer with mock configuration
	producer, err := NewProducer(MockProducerConfig)
	assert.NoError(t, err)
	assert.NotNil(t, producer)

	// Test the Close method
	err = producer.Close()
	assert.NoError(t, err)
}

func TestKafkaProducer_SafeClose(t *testing.T) {
	// Create a KafkaProducer with mock configuration
	producer, err := NewProducer(MockProducerConfig)
	assert.NoError(t, err)
	assert.NotNil(t, producer)

	// Test the SafeClose method
	err = producer.SafeClose()
	assert.NoError(t, err)
}
