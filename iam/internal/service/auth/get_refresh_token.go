package auth

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/utils/jwt_tokens"
)

func (s *service) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := jwt_tokens.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	_, err = s.authRepository.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	newRefreshToken, _, err := jwt_tokens.GenerateRefreshToken(claims.UserId, claims.Username, claims.Role, s.refreshTokenTtl)
	if err != nil {
		return "", err
	}

	// Should be wrapped in tx but not worth it, if it is an issue ill change that
	err = s.authRepository.DeleteRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	_, err = s.authRepository.CreateRefreshToken(ctx, claims.UserId, newRefreshToken, s.refreshTokenTtl)
	if err != nil {
		return "", err
	}

	return newRefreshToken, nil
}
