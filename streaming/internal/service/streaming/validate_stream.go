package streaming

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/streaming/internal/model"
)

func (s *service) ValidateStream(ctx context.Context, streamKey string) error {
	user, err := s.iamClient.GetUserViaStreamKey(ctx, streamKey)
	if err != nil {
		return model.ErrInvalidStreamKey
	}

	err = s.streamingProducerService.ProduceStreamStarted(ctx, model.StreamStartedEvent{
		UserId: user.UserId,
	})
	if err != nil {
		return err
	}

	return nil
}
