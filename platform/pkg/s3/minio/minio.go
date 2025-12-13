package minio

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

type client struct {
	db     *minio.Client
	logger Logger
	url    string
}

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

func NewMinioClient(logger Logger, db *minio.Client, url string) *client {
	return &client{
		db:     db,
		logger: logger,
		url:    url,
	}
}

func (c *client) UploadFile(ctx context.Context, bucket string, filename string, reader io.Reader, size int64, contentType string) error {
	_, err := c.db.PutObject(ctx, bucket, filename, reader, size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		c.logger.Error(ctx, "could not upload file", zap.String("name", filename), zap.Error(err))
		return err
	}

	return nil
}

func (c *client) GetFile(ctx context.Context, bucket string, filename string) (*minio.Object, error) {
	obj, err := c.db.GetObject(ctx, bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		c.logger.Error(ctx, "could not get file", zap.String("name", filename), zap.Error(err))
		return nil, err
	}

	return obj, nil
}

func (c *client) GetFileUrl(bucket string, filename string) string {
	return fmt.Sprintf("%s/%s/%s", c.url, bucket, filename)
}
