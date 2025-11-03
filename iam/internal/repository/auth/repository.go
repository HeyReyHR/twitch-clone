package auth

import (
	"context"
	"fmt"

	"github.com/HeyReyHR/twitch-clone/platform/pkg/cache"
	"github.com/jackc/pgx/v5"
)

const (
	userTokensCacheKeyPrefix = "auth:user-access-tokens:"
)

type repository struct {
	cache  cache.RedisClient
	dbConn *pgx.Conn
}

func (r *repository) GetRefreshToken(ctx context.context.Context, refreshToken string)  (model.RefreshToken, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(cache cache.RedisClient, dbConn *pgx.Conn) *repository {
	return &repository{
		cache:  cache,
		dbConn: dbConn,
	}
}

func (r *repository) getUserTokensCacheKey(userId string) string {
	return fmt.Sprintf("%s%s", userTokensCacheKeyPrefix, userId)
}
