package service

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
)

type UserService interface {
	Register(ctx context.Context, email string, username string, role model.Role, password string, passwordConfirmation string) (string, error)
	Get(ctx context.Context, userId string) (model.User, error)
}
