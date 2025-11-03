package auth

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/repository/convert"
	"github.com/HeyReyHR/twitch-clone/iam/internal/utils/jwt_tokens"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := jwt_tokens.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	accessToken, err := s.authRepository.GetAccessToken(ctx, claims.UserId)
	if err != nil {
		return "", err
	}
	if accessToken != "" {
		return accessToken, nil
	}

	user, err := s.userRepository.Get(ctx, claims.UserId)
	if err != nil {
		return "", err
	}

	_, err = s.authRepository.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, _, err = jwt_tokens.GenerateAccessToken(convert.RepoToServiceUser(user), s.accessTokenTtl)

	err = s.authRepository.CreateAccessToken(ctx, claims.UserId, accessToken, s.accessTokenTtl)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
