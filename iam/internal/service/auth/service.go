package auth

import (
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/repository"
)

type service struct {
	authRepository  repository.AuthRepository
	userRepository  repository.UserRepository
	accessTokenTtl  time.Duration
	refreshTokenTtl time.Duration
}

func NewService(authRepository repository.AuthRepository, userRepository repository.UserRepository, accessTokenTtl time.Duration, refreshTokenTtl time.Duration) *service {
	return &service{
		authRepository:  authRepository,
		userRepository:  userRepository,
		accessTokenTtl:  accessTokenTtl,
		refreshTokenTtl: refreshTokenTtl,
	}
}
