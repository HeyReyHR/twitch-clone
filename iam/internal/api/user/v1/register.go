package v1

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/convert"
	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
)

func (a *api) Register(ctx context.Context, r *userV1.RegisterRequest) (*userV1.RegisterResponse, error) {
	userId, err := a.userService.Register(ctx, r.GetEmail(), r.GetUsername(), convert.RoleApiToService(r.GetRole()), r.GetPassword(), r.GetPasswordConfirmation())
	if err != nil {
		return nil, err
	}

	return &userV1.RegisterResponse{
		UserId: userId,
	}, nil
}
