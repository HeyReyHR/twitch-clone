package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type streamingHTTPEnvConfig struct {
	Host string `env:"HTTP_HOST,required"`
	Port string `env:"HTTP_PORT,required"`
}

type streamingHTTPConfig struct {
	raw streamingHTTPEnvConfig
}

func NewStreamingHTTPConfig() (*streamingHTTPConfig, error) {
	var raw streamingHTTPEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}
	return &streamingHTTPConfig{raw: raw}, nil
}

func (cfg *streamingHTTPConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}
