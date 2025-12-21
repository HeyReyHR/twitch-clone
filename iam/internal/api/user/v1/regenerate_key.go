package v1

import (
	"context"

	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
)

func (a *api) RegenerateStreamKey(ctx context.Context, r *userV1.RegenerateStreamKeyRequest) (*userV1.RegenerateStreamKeyResponse, error) {
	key, err := a.userService.RegenerateKey(ctx, r.GetUserId())
	if err != nil {
		return nil, err
	}

	return &userV1.RegenerateStreamKeyResponse{
		StreamKey: key,
	}, nil
}
