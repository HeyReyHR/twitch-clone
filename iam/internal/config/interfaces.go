package config

import "time"

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type ServiceConfig interface {
	Address() string
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
	MigrationsDir() string
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

type JWTTokensConfig interface {
	AccessTokenTTL() time.Duration
	RefreshTokenTTL() time.Duration
}

type PasswordConfig interface {
	PasswordSalt() string
}
