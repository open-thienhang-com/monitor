package kafka

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	ERROR_TYPE_PARSE_MSG       = "parse_msg_error"
	ERROR_TYPE_HANDLE_MSG      = "handle_msg_error"
	ERROR_TYPE_PUSH_DEADLETTER = "push_dead_letter_error"
)

var doOnce sync.Once
var (
	ProcessDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "app",
		Subsystem: "kafka",
		Name:      "kafka_msg_consuming_duration_ms",
		Help:      "Kafka consumer message processing duration in ms",
	}, []string{"topic"})

	MessageSize = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "app",
		Subsystem: "kafka",
		Name:      "kafka_msg_consuming_size_bytes",
		Help:      "Kafka message size in bytes",
	}, []string{"topic"})

	ErrorCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "app",
		Subsystem: "kafka",
		Name:      "kafka_msg_consuming_error_count",
		Help:      "Kafka consuming errors count",
	}, []string{"topic", "err_type"})

	SuccessCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "app",
		Subsystem: "kafka",
		Name:      "kafka_msg_consuming_success_count",
		Help:      "Kafka consuming success count",
	}, []string{"topic"})
)

func init() {
	doOnce.Do(func() {
		prometheus.MustRegister(
			ProcessDuration,
			MessageSize,
			ErrorCount,
			SuccessCount,
		)
	})
}

func LogError(topic string, err error, errorType string) {
	ErrorCount.WithLabelValues(topic, errorType).Inc()
}

func LogSuccess(topic string) {
	SuccessCount.WithLabelValues(topic).Inc()
}

func LogConsumingDuration(topic string, duration float64) {
	ProcessDuration.WithLabelValues(topic).Observe(duration)
}

func LogMessageSize(topic string, size float64) {
	MessageSize.WithLabelValues(topic).Observe(size)
}
