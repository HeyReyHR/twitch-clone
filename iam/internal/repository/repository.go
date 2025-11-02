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
	Create(ctx context.Context, userId string, refreshToken string, tokenTtl time.Duration) error
	Get(ctx context.Context, userId string) (string, error)
}
