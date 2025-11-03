package v1

import (
	"context"

	authV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/auth/v1"
)

func (a *api) GetAccessToken(ctx context.Context, r *authV1.GetAccessTokenRequest) (*authV1.GetAccessTokenResponse, error) {
	accessToken, err := a.authService.GetAccessToken(ctx, r.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &authV1.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
