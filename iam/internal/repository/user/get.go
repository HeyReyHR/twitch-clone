package user

import (
	"context"
	"errors"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) Get(ctx context.Context, userId string) (repoModel.User, error) {
	var user repoModel.User

	err := r.dbConn.QueryRow(ctx, "SELECT user_id, username, email, role, created_at, updated_at FROM users WHERE user_id = $1", userId).Scan(
		&user.UserId,
		&user.Username,
		&user.UserId,
		&user.Role,
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
