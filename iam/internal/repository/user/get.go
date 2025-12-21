package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/logger"
)

func (r *repository) Get(ctx context.Context, userId string) (repoModel.User, error) {
	var user repoModel.User

	err := r.dbConn.QueryRow(ctx, "SELECT user_id, username, email, password_hash, role, avatar_url, is_streaming, stream_key, created_at, updated_at FROM users WHERE user_id = $1", userId).Scan(
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

		logger.Debug(ctx, "postgres error", zap.Error(err))
		return repoModel.User{}, model.ErrDbScanFailed
	}

	return user, nil
}
