package auth

import (
	"context"
	"errors"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	redigo "github.com/gomodule/redigo/redis"
)

func (r *repository) GetAccessToken(ctx context.Context, userId string) (string, error) {
	accessToken, err := r.cache.Get(ctx, r.getUserTokensCacheKey(userId))
	if err != nil {
		if errors.Is(err, redigo.ErrNil) {
			return "", model.ErrRedisCacheNotFound
		}
		return "", model.ErrRedisGetCacheFailed
	}

	return string(accessToken), nil
}
