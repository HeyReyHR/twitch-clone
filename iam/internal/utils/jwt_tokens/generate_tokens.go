package jwt_tokens

import (
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/config"
	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenPair(user model.User, accessTokenTtl time.Duration, refreshTokenTtl time.Duration) (*model.TokenPair, error) {
	accessToken, accessExpiresAt, err := GenerateAccessToken(user, accessTokenTtl)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := GenerateRefreshToken(user, refreshTokenTtl)
	if err != nil {
		return nil, err
	}

	return &model.TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessExpiresAt,
		RefreshTokenExpiresAt: refreshExpiresAt,
	}, nil
}

func GenerateAccessToken(user model.User, accessTokenTtl time.Duration) (string, time.Time, error) {
	expiresAt := time.Now().Add(accessTokenTtl)

	claims := jwt.MapClaims{
		"user_id":  user.UserId,
		"username": user.Username,
		"role":     user.Role,
		"exp":      expiresAt.Unix(),
		"iat":      time.Now().Unix(),
		"type":     "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig().JWTTokens.AccessTokenSecret()))
	if err != nil {
		return "", time.Time{}, model.ErrGenerateTokenFailed
	}

	return tokenString, expiresAt, nil
}

func GenerateRefreshToken(user model.User, refreshTokenTtl time.Duration) (string, time.Time, error) {
	expiresAt := time.Now().Add(refreshTokenTtl)

	claims := jwt.MapClaims{
		"user_id":  user.UserId,
		"username": user.Username,
		"role":     user.Role,
		"exp":      expiresAt.Unix(),
		"iat":      time.Now().Unix(),
		"type":     "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig().JWTTokens.RefreshTokenSecret()))
	if err != nil {
		return "", time.Time{}, model.ErrGenerateTokenFailed
	}

	return tokenString, expiresAt, nil
}
