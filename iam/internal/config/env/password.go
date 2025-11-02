package env

import (
	"github.com/caarlos0/env/v11"
)

type passwordEnvConfig struct {
	PasswordSalt string `env:"PASSWORD_SALT,required"`
}

type passwordConfig struct {
	raw passwordEnvConfig
}

func NewPasswordConfig() (*passwordConfig, error) {
	var raw passwordEnvConfig
	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}

	return &passwordConfig{raw: raw}, nil
}

func (cfg *passwordConfig) PasswordSalt() string {
	return cfg.raw.PasswordSalt
}
