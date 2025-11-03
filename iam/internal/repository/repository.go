package repository

import (
	"context"
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
)

type UserRepository interface {
	Create(ctx context.Context, email, username string, role model.Role, passwordHash string) (string, error)
	Get(ctx context.Context, userId string) (model.User, error)
	GetViaEmail(ctx context.Context, email string) (model.User, error)
	GetViaUsername(ctx context.Context, username string) (model.User, error)
}

type AuthRepository interface {
	CreateRefreshToken(ctx context.Context, userId, refreshToken string, expiresAt time.Duration) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (model.RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, refreshToken string) error
}
