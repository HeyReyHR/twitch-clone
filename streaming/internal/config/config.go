package config

import (
	"os"

	"github.com/HeyReyHR/twitch-clone/streaming/internal/config/env"
	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	IamGRPC               ServiceConfig
	StreamingHttp         ServiceConfig
	Kafka                 KafkaConfig
	Logger                LoggerConfig
	StreamStartedProducer ProducerConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	iamGRPCCfg, err := env.NewIamClientGRPCConfig()
	if err != nil {
		return err
	}

	kafkaCfg, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	streamingHTTPCfg, err := env.NewStreamingHTTPConfig()
	if err != nil {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	streamStartedProducerCfg, err := env.NewStreamStartedProducerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		IamGRPC:               iamGRPCCfg,
		Kafka:                 kafkaCfg,
		Logger:                loggerCfg,
		StreamingHttp:         streamingHTTPCfg,
		StreamStartedProducer: streamStartedProducerCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
