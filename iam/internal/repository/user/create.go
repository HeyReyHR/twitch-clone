package user

import (
	"context"
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
	"github.com/google/uuid"
)

func (r *repository) Create(ctx context.Context, email string, username string, role repoModel.Role, passwordHash string) (string, error) {
	userId := uuid.NewString()
	currentTime := time.Now()
	_, err := r.dbConn.Exec(ctx, "INSERT INTO users (user_id, username, email, role, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4)", userId, username, email, passwordHash, currentTime, currentTime)
	if err != nil {
		return "", model.ErrCreateUserFailed
	}

	return userId, nil
}
