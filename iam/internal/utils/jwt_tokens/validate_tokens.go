package jwt_tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/HeyReyHR/twitch-clone/iam/internal/config"
	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
)

func ValidateRefreshToken(tokenString string) (*model.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrMalformedToken
		}
		return []byte(config.AppConfig().JWTTokens.RefreshTokenSecret()), nil
	})
	if err != nil || !token.Valid {
		return nil, model.ErrMalformedToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, model.ErrMalformedToken
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return nil, model.ErrMalformedToken
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, model.ErrMalformedToken
	}

	expiresAt := time.Unix(int64(exp), 0)
	if time.Now().After(expiresAt) {
		return nil, model.ErrMalformedToken
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		return nil, model.ErrMalformedToken
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, model.ErrMalformedToken
	}

	roleToken, ok := claims["role"].(string)
	if !ok {
		return nil, model.ErrMalformedToken
	}

	role := model.Role(roleToken)
	if role != model.USER && role != model.ADMIN {
		return nil, model.ErrMalformedToken
	}

	return &model.Claims{
		UserId:   userId,
		Username: username,
		Role:     role,
	}, nil
}
