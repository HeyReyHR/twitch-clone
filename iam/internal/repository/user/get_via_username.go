package user

import (
	"context"
	"errors"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) GetViaUsername(ctx context.Context, username string) (repoModel.User, error) {
	var user repoModel.User

	err := r.dbConn.QueryRow(ctx, "SELECT user_id, username, email, role, password_hash, created_at, updated_at FROM users WHERE username = $1", username).Scan(
		&user.UserId,
		&user.Username,
		&user.UserId,
		&user.Role,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repoModel.User{}, model.ErrUserNotFound
		}
		return repoModel.User{}, model.ErrUserScanFailed
	}

	return user, nil

}
