package auth

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/utils/jwt_tokens"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := jwt_tokens.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	_, err = s.authRepository.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, _, err := jwt_tokens.GenerateAccessToken(claims.UserId, claims.Username, claims.Role, s.accessTokenTtl)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
