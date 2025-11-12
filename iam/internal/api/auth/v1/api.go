package v1

import (
	authv3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"

	"github.com/HeyReyHR/twitch-clone/iam/internal/service"
	authV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/auth/v1"
)

type api struct {
	authV1.UnimplementedAuthServiceServer
	authv3.UnimplementedAuthorizationServer

	authService service.AuthService
}

func NewApi(authService service.AuthService) *api {
	return &api{
		authService: authService,
	}
}
