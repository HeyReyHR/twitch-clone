package minio

import (
	"context"
	"fmt"
	"mime"
	"path/filepath"

	"github.com/google/uuid"
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

func (c *client) UploadFile(ctx context.Context, bucket string, path string) {
	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)
	// TODO userId/avatar_uuid.png, maybe change uuid to time.Now().Unix()
	name := fmt.Sprintf("avatar_%s.%s", uuid.NewString(), ext)

	_, err := c.db.FPutObject(ctx, bucket, name, path, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		c.logger.Error(ctx, "could not upload file", zap.String("name", name), zap.Error(err))
	}
}
