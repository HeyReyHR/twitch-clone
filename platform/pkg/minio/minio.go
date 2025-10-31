package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

type client struct {
	db     *minio.Client
	logger Logger
}

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

func NewMinioClient(logger Logger, db *minio.Client) *client {
	return &client{
		db:     db,
		logger: logger,
	}
}
