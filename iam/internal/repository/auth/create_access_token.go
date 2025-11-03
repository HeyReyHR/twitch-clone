package auth

import (
	"context"
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
)

func (r *repository) CreateAccessToken(ctx context.Context, userId string, accessToken string, tokenTtl time.Duration) error {
	if err := r.cache.SAdd(ctx, r.getUserTokensCacheKey(userId), accessToken); err != nil {
		return model.ErrRedisEditCacheDbFailed
	}
	if err := r.cache.Expire(ctx, r.getUserTokensCacheKey(userId), tokenTtl); err != nil {
		return model.ErrRedisEditCacheDbFailed
	}

	return nil
}
