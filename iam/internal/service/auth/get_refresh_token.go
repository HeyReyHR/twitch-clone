package auth

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/repository/convert"
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

	user, err := s.userRepository.Get(ctx, claims.UserId)
	if err != nil {
		return "", err
	}

	refreshToken, _, err = jwt_tokens.GenerateRefreshToken(convert.RepoToServiceUser(user), s.refreshTokenTtl)
	_, err = s.authRepository.CreateRefreshToken(ctx, claims.UserId, refreshToken, s.refreshTokenTtl)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
