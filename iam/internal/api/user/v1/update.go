package v1

import (
	"context"

	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
)

func (a *api) Update(ctx context.Context, r *userV1.UpdateRequest) (*userV1.UpdateResponse, error) {
	err := a.userService.Update(ctx, r.UserId, r.Username, r.Email, r.Avatar, r.AvatarContentType, r.IsStreaming)
	if err != nil {
		return nil, err
	}

	return &userV1.UpdateResponse{}, nil
}
