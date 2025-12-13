package config

import (
	"time"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

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

type JWTTokensConfig interface {
	AccessTokenTTL() time.Duration
	RefreshTokenTTL() time.Duration
	AccessTokenSecret() string
	RefreshTokenSecret() string
}

type PasswordConfig interface {
	PasswordSalt() string
}

type MinioConfig interface {
	Endpoint() string
	PublicUrl() string
	Credentials() *credentials.Credentials
	AvatarBucket() string
}
