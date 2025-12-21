package v1

import (
	"context"

	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
	"github.com/HeyReyHR/twitch-clone/streaming/internal/model"
)

type client struct {
	generatedClient userV1.UserServiceClient
}

func NewUserClient(generatedClient userV1.UserServiceClient) *client {
	return &client{
		generatedClient: generatedClient,
	}
}

func (c *client) GetUserViaStreamKey(ctx context.Context, streamKey string) (*model.User, error) {
	resp, err := c.generatedClient.GetUserViaStreamKey(ctx, &userV1.GetUserViaStreamKeyRequest{
		StreamKey: streamKey,
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		UserId:      resp.GetUser().GetUserId(),
		Username:    resp.GetUser().GetUsername(),
		Email:       resp.GetUser().GetEmail(),
		IsStreaming: resp.GetUser().GetIsStreaming(),
		StreamKey:   resp.GetUser().GetStreamKey(),
	}, nil
}
