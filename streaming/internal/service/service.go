package service

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/streaming/internal/model"
)

type StreamingService interface {
	ValidateStream(ctx context.Context, streamKey string) error
}

type StreamingProducerService interface {
	ProduceStreamStarted(ctx context.Context, event model.StreamStartedEvent) error
}
