package s3

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

type MinioClient interface {
	UploadFile(ctx context.Context, bucket string, filename string, reader io.Reader, size int64, contentType string) error
	GetFile(ctx context.Context, bucket string, filename string) (*minio.Object, error)
	GetFileUrl(bucket string, filename string) string
}
