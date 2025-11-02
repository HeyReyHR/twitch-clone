package v1

import (
	"github.com/HeyReyHR/twitch-clone/iam/internal/service"
	userV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/user/v1"
)

type api struct {
	userV1.UnimplementedUserServiceServer

	userService service.UserService
}

func NewApi(userService service.UserService) *api {
	return &api{
		userService: userService,
	}
}
