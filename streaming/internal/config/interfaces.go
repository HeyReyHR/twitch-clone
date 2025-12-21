package config

import "github.com/IBM/sarama"

type ServiceConfig interface {
	Address() string
}

type KafkaConfig interface {
	Brokers() []string
}

type ProducerConfig interface {
	Topic() string
	Config() *sarama.Config
}

type LoggerConfig interface {
	Level() string
	AsJson() bool
}
