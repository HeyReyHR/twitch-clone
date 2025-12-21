package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	commonV1 "github.com/HeyReyHR/twitch-clone/shared/pkg/proto/common/v1"
)

func RoleApiToService(role commonV1.Role) model.Role {
	switch role {
	case commonV1.Role_USER:
		return model.USER
	case commonV1.Role_ADMIN:
		return model.ADMIN
	default:
		return model.UNKNOWN
	}
}

func RoleServiceToApi(role model.Role) commonV1.Role {
	switch role {
	case model.USER:
		return commonV1.Role_USER
	case model.ADMIN:
		return commonV1.Role_ADMIN
	default:
		return commonV1.Role_UNKNOWN
	}
}

func UserServiceToApi(user *model.User) *commonV1.User {
	return &commonV1.User{
		UserId:      user.UserId,
		Username:    user.Username,
		Email:       user.Email,
		AvatarUrl:   user.AvatarUrl,
		IsStreaming: user.IsStreaming,
		StreamKey:   user.StreamKey,
		Role:        RoleServiceToApi(user.Role),
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
	}
}
