package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type streamStartedProducerEnvConfig struct {
	TopicName string `env:"STREAM_STARTED_TOPIC_NAME,required"`
}

type streamStartedProducerConfig struct {
	raw streamStartedProducerEnvConfig
}

func NewStreamStartedProducerConfig() (*streamStartedProducerConfig, error) {
	var raw streamStartedProducerEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &streamStartedProducerConfig{raw: raw}, nil
}

func (cfg *streamStartedProducerConfig) Topic() string {
	return cfg.raw.TopicName
}

func (cfg *streamStartedProducerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Producer.Return.Successes = true

	return config
}
