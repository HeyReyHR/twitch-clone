package user

import (
	"github.com/HeyReyHR/twitch-clone/iam/internal/repository"
)

type service struct {
	repository repository.UserRepository
}

func NewService(repository repository.UserRepository) *service {
	return &service{
		repository: repository,
	}
}
