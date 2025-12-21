package user

import (
	"context"
	"errors"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) GetViaStreamKey(ctx context.Context, streamKey string) (repoModel.User, error) {
	var user repoModel.User
	err := r.dbConn.QueryRow(ctx, "SELECT user_id, username, email, password_hash, role, avatar_url, is_streaming, stream_key, created_at, updated_at FROM users WHERE stream_key = $1", streamKey).Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.AvatarUrl,
		&user.IsStreaming,
		&user.StreamKey,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repoModel.User{}, model.ErrDbEntityNotFound
		}
		return repoModel.User{}, model.ErrDbScanFailed
	}

	return user, nil
}
