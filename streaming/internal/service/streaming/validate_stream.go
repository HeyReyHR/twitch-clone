package streaming

import (
	"context"
	"fmt"

	"github.com/HeyReyHR/twitch-clone/streaming/internal/model"
)

func (s *service) ValidateStream(ctx context.Context, streamKey string) error {
	user, err := s.iamClient.GetUserViaStreamKey(ctx, streamKey)
	if err != nil {

		fmt.Println("user", err)
		return model.ErrInvalidStreamKey
	}

	err = s.streamingProducerService.ProduceStreamStarted(ctx, model.StreamStartedEvent{
		UserId: user.UserId,
	})
	fmt.Println("prod", err)
	if err != nil {
		return err
	}

	return nil
}
