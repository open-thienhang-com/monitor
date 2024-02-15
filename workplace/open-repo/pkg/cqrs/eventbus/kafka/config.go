package kafka

import "mono.thienhang.com/pkg/eventbus"

type Config struct {
	eventbus.Config
	consumer ConsumerConfig
	producer ProducerConfig
}

type ConsumerConfig struct {
	Brokers           []string `mapstructure:"brokers"`
	GroupID           string   `mapstructure:"groupId"`
	Topic             string   `mapstructure:"topic"`
	Mock              bool     `mapstructure:"mock"`
	PoolSize          int      `mapstructure:"poolSize"`
	ReleaseTimeoutSec int      `mapstructure:"releaseTimeoutSec"`
	DeadLetterTopic   string   `mapstructure:"deadLetterTopic"`
	Replication       int      `mapstructure:"replication"`
	Partition         int      `mapstructure:"partition"`
}

type ProducerConfig struct {
	Brokers           []string `mapstructure:"brokers"`
	Topic             string   `mapstructure:"topic"`
	Mock              bool     `mapstructure:"mock"`
	Replication       int      `mapstructure:"replication"`
	Partition         int      `mapstructure:"partition"`
	BufferSize        int      `mapstructure:"bufferSize"`
	PoolSize          int      `mapstructure:"poolSize"`
	ReleaseTimeoutSec int      `mapstructure:"releaseTimeoutSec"`
}
