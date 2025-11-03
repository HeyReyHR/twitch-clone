package v1

import (
	"context"

	authV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/auth/v1"
)

func (a *api) GetRefreshToken(ctx context.Context, r *authV1.GetRefreshTokenRequest) (*authV1.GetRefreshTokenResponse, error) {
	refreshToken, err := a.authService.GetRefreshToken(ctx, r.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &authV1.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}
