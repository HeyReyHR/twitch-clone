package auth

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
)

func (r *repository) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	_, err := r.dbConn.Exec(ctx, "DELETE FROM refresh_tokens WHERE refresh_token = $1", refreshToken)
	if err != nil {
		return model.ErrDeletionFailed
	}
	return nil
}
