package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioEnvConfig struct {
	Host         string `env:"MINIO_HOST,required"`
	Port         string `env:"MINIO_PORT,required"`
	AccessKey    string `env:"MINIO_ACCESS_KEY,required"`
	SecretKey    string `env:"MINIO_SECRET_KEY,required"`
	AvatarBucket string `env:"AVATAR_BUCKET,required"`
}
type minioConfig struct {
	raw minioEnvConfig
}

func NewMinioConfig() (*minioConfig, error) {
	var raw minioEnvConfig
	err := env.Parse(&raw)
	if err != nil {
		return nil, err
	}

	return &minioConfig{raw: raw}, nil
}

func (cfg *minioConfig) Endpoint() string {
	return fmt.Sprintf(
		"%s:%s",
		cfg.raw.Host,
		cfg.raw.Port,
	)
}

func (cfg *minioConfig) PublicUrl() string {
	return fmt.Sprintf("http://%s:%s", cfg.raw.Host, cfg.raw.Port)
}

func (cfg *minioConfig) Credentials() *credentials.Credentials {
	return credentials.NewStaticV4(cfg.raw.AccessKey, cfg.raw.SecretKey, "")
}

func (cfg *minioConfig) AvatarBucket() string {
	return cfg.raw.AvatarBucket
}
