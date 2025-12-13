package user

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
)

func (r *repository) Update(ctx context.Context, userId string, params repoModel.UpdateParams) error {
	var setClauses []string
	var args []any
	var argIndex uint8 = 1

	if params.Username != nil {
		setClauses = append(setClauses, fmt.Sprintf("username = $%d", argIndex))
		args = append(args, *params.Username)
		argIndex++
	}
	if params.Email != nil {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argIndex))
		args = append(args, *params.Email)
		argIndex++
	}
	if params.Role != nil {
		setClauses = append(setClauses, fmt.Sprintf("role = $%d", argIndex))
		args = append(args, *params.Role)
		argIndex++
	}
	if params.AvatarUrl != nil {
		setClauses = append(setClauses, fmt.Sprintf("avatar_url = $%d", argIndex))
		args = append(args, *params.AvatarUrl)
		argIndex++
	}
	if params.IsStreaming != nil {
		setClauses = append(setClauses, fmt.Sprintf("is_streaming = $%d", argIndex))
		args = append(args, *params.IsStreaming)
		argIndex++
	}

	if len(setClauses) == 0 {
		return nil
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", argIndex))
	args = append(args, time.Now())
	argIndex++

	args = append(args, userId)

	query := fmt.Sprintf("UPDATE users SET %s WHERE user_id = $%d", strings.Join(setClauses, ", "), argIndex)

	_, err := r.dbConn.Exec(ctx, query, args...)
	if err != nil {
		return model.ErrUpdateDbEntityFailed
	}

	return nil
}
