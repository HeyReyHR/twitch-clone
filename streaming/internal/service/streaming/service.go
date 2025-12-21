package streaming

import (
	"github.com/HeyReyHR/twitch-clone/streaming/internal/client"
	def "github.com/HeyReyHR/twitch-clone/streaming/internal/service"
)

type service struct {
	iamClient                client.IamClient
	streamingProducerService def.StreamingProducerService
}

func NewService(iamClient client.IamClient, streamingProducerService def.StreamingProducerService) *service {
	return &service{
		iamClient:                iamClient,
		streamingProducerService: streamingProducerService,
	}
}
