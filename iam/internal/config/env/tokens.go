package env

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type jwtTokensEnvConfig struct {
	AccessTokenTTL     time.Duration `env:"ACCESS_TOKEN_TTL,required"`
	RefreshTokenTTL    time.Duration `env:"REFRESH_TOKEN_TTL,required"`
	AccessTokenSecret  string        `env:"ACCESS_TOKEN_SECRET,required"`
	RefreshTokenSecret string        `env:"REFRESH_TOKEN_SECRET,required"`
}

type jwtTokensConfig struct {
	raw jwtTokensEnvConfig
}

func NewJWTTokensConfig() (*jwtTokensConfig, error) {
	var raw jwtTokensEnvConfig
	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}

	return &jwtTokensConfig{raw: raw}, nil
}

func (cfg *jwtTokensConfig) AccessTokenTTL() time.Duration {
	return cfg.raw.AccessTokenTTL
}

func (cfg *jwtTokensConfig) RefreshTokenTTL() time.Duration {
	return cfg.raw.RefreshTokenTTL
}

func (cfg *jwtTokensConfig) AccessTokenSecret() string {
	return cfg.raw.AccessTokenSecret
}

func (cfg *jwtTokensConfig) RefreshTokenSecret() string {
	return cfg.raw.RefreshTokenSecret
}
