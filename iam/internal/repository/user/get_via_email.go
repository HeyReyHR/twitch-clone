package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
)

func (r *repository) GetViaEmail(ctx context.Context, email string) (repoModel.User, error) {
	var user repoModel.User

	err := r.dbConn.QueryRow(ctx, "SELECT user_id, username, email, role, password_hash, created_at, updated_at FROM users WHERE email = $1", email).Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
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
