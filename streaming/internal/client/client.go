package client

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/streaming/internal/model"
)

type IamClient interface {
	GetUserViaStreamKey(ctx context.Context, streamKey string) (model.User, error)
}
