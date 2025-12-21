package v1

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/convert"
	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
)

func (a *api) GetUserViaStreamKey(ctx context.Context, r *userV1.GetUserViaStreamKeyRequest) (*userV1.GetUserViaStreamKeyResponse, error) {
	user, err := a.userService.GetViaStreamKey(ctx, r.GetStreamKey())
	if err != nil {
		return nil, err
	}

	return &userV1.GetUserViaStreamKeyResponse{
		User: convert.UserServiceToApi(user),
	}, nil
}
