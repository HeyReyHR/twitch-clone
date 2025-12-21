package v1

import "github.com/HeyReyHR/twitch-clone/streaming/internal/service"

type api struct {
	service service.StreamingService
}

func NewApi(service service.StreamingService) *api {
	return &api{
		service: service,
	}
}
