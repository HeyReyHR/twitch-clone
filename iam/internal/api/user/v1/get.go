package v1

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/convert"
	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
)

func (a *api) GetUser(ctx context.Context, r *userV1.GetUserRequest) (*userV1.GetUserResponse, error) {
	user, err := a.userService.Get(ctx, r.GetUserId())
	if err != nil {
		return nil, err
	}

	return &userV1.GetUserResponse{
		User: convert.UserServiceToApi(user),
	}, nil
}
