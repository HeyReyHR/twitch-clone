package streaming_producer

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/platform/pkg/kafka"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/logger"
	eventsV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/events/v1"
	"github.com/HeyReyHR/twitch-clone/streaming/internal/model"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type service struct {
	streamStartedProducer kafka.Producer
}

func NewService(streamStartedProducer kafka.Producer) *service {
	return &service{
		streamStartedProducer: streamStartedProducer,
	}
}

func (p *service) ProduceStreamStarted(ctx context.Context, event model.StreamStartedEvent) error {
	msg := &eventsV1.StreamStarted{
		UserId: event.UserId,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal StreamStarted", zap.Error(err))
		return err
	}

	err = p.streamStartedProducer.Send(ctx, []byte(event.EventId), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish StreamStarted", zap.Error(err))
		return err
	}

	return nil
}
