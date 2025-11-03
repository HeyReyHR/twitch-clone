package repository

import (
	"context"
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
)

type UserRepository interface {
	Create(ctx context.Context, email string, username string, role model.Role, passwordHash string) (string, error)
	Get(ctx context.Context, userId string) (model.User, error)
	GetViaEmail(ctx context.Context, email string) (model.User, error)
	GetViaUsername(ctx context.Context, username string) (model.User, error)
}

type AuthRepository interface {
	CreateAccessToken(ctx context.Context, userId string, accessToken string, tokenTtl time.Duration) error
	CreateRefreshToken(ctx context.Context, userId string, refreshToken string, expiresAt time.Duration) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (model.RefreshToken, error)
	GetAccessToken(ctx context.Context, userId string) (string, error)
}
