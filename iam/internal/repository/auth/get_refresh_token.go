package auth

import (
	"context"
	"errors"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) Get(ctx context.Context, refreshToken string) (repoModel.RefreshToken, error) {
	var token repoModel.RefreshToken

	err := r.dbConn.QueryRow(ctx, "SELECT id, user_id, refresh_token, created_at, expires_at FROM refresh_tokens WHERE refresh_token = $1", refreshToken).Scan(
		&token.Id,
		&token.UserId,
		&token.RefreshToken,
		&token.CreatedAt,
		&token.ExpiresAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repoModel.RefreshToken{}, model.ErrDbEntityNotFound
		}
		return repoModel.RefreshToken{}, model.ErrDbScanFailed
	}

	return token, nil

}
