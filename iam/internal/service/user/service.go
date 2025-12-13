package user

import (
	"github.com/HeyReyHR/twitch-clone/iam/internal/repository"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/s3"
)

type service struct {
	repository repository.UserRepository
	minio      s3.MinioClient
}

func NewService(repository repository.UserRepository, minio s3.MinioClient) *service {
	return &service{
		repository: repository,
		minio:      minio,
	}
}
