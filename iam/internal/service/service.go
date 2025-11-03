package service

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
)

type UserService interface {
	Register(ctx context.Context, email, username string, role model.Role, password, passwordConfirmation string) (string, error)
	Get(ctx context.Context, userId string) (*model.User, error)
}

type AuthService interface {
	Login(ctx context.Context, email, password string) (*model.TokenPair, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
}
