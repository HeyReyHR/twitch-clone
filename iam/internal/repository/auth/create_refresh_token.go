package auth

import (
	"context"
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	"github.com/google/uuid"
)

func (r *repository) CreateRefreshToken(ctx context.Context, userId, refreshToken string, expiresAt time.Duration) (string, error) {
	tokenId := uuid.NewString()

	currentTime := time.Now()
	expiration := currentTime.Add(expiresAt)

	_, err := r.dbConn.Exec(ctx, "INSERT INTO refresh_tokens (id, user_id, refresh_token, created_at, expires_at) VALUES ($1, $2, $3, $4, $5)", tokenId, userId, refreshToken, currentTime, expiration)
	if err != nil {
		return "", model.ErrCreateDbEntityFailed
	}

	return tokenId, nil
}
