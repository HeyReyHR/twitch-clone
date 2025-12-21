package app

import (
	"context"
	"fmt"

	"github.com/HeyReyHR/twitch-clone/platform/pkg/closer"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/kafka"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/kafka/producer"
	"github.com/HeyReyHR/twitch-clone/platform/pkg/logger"
	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
	"github.com/HeyReyHR/twitch-clone/streaming/internal/api"
	streamingV1API "github.com/HeyReyHR/twitch-clone/streaming/internal/api/streaming/v1"
	"github.com/HeyReyHR/twitch-clone/streaming/internal/client"
	iamClientV1 "github.com/HeyReyHR/twitch-clone/streaming/internal/client/iam/v1"
	"github.com/HeyReyHR/twitch-clone/streaming/internal/config"
	"github.com/HeyReyHR/twitch-clone/streaming/internal/service"
	streamingProducer "github.com/HeyReyHR/twitch-clone/streaming/internal/service/producer/streaming_producer"
	streamingService "github.com/HeyReyHR/twitch-clone/streaming/internal/service/streaming"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type diContainer struct {
	streamingV1API api.StreamingHandlerV1

	iamClient client.IamClient

	streamingProducerService service.StreamingProducerService
	streamingService         service.StreamingService

	syncProducer          sarama.SyncProducer
	streamStartedProducer kafka.Producer
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) StreamingV1API(ctx context.Context) api.StreamingHandlerV1 {
	if d.streamingV1API == nil {
		d.streamingV1API = streamingV1API.NewApi(d.StreamingService(ctx))
	}

	return d.streamingV1API
}

func (d *diContainer) StreamingService(ctx context.Context) service.StreamingService {
	if d.streamingService == nil {
		d.streamingService = streamingService.NewService(d.IamClient(ctx), d.StreamingProducerService())
	}

	return d.streamingService
}

func (d *diContainer) StreamingProducerService() service.StreamingProducerService {
	if d.streamingProducerService == nil {
		d.streamingProducerService = streamingProducer.NewService(d.StreamStartedProducer())
	}

	return d.streamingProducerService
}

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().StreamStartedProducer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create sync producer: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka sync producer", func(ctx context.Context) error {
			return p.Close()
		})

		d.syncProducer = p
	}

	return d.syncProducer
}

func (d *diContainer) StreamStartedProducer() kafka.Producer {
	if d.streamStartedProducer == nil {
		d.streamStartedProducer = producer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().StreamStartedProducer.Topic(),
			logger.Logger())
	}

	return d.streamStartedProducer
}

func (d *diContainer) IamClient(ctx context.Context) client.IamClient {
	if d.iamClient == nil {
		connIam, err := grpc.NewClient(
			config.AppConfig().IamGRPC.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Error(ctx, "❌ failed to connect to iam service", zap.Error(err))
			return nil
		}

		iam := userV1.NewUserServiceClient(connIam)

		d.iamClient = iamClientV1.NewUserClient(iam)
		closer.AddNamed("Iam gRPC client", func(ctx context.Context) error {
			return connIam.Close()
		})
	}
	return d.iamClient
}
