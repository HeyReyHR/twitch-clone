package convert

import (
	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
)

func ServiceToRepoRole(role model.Role) repoModel.Role {
	switch role {
	case model.USER:
		return repoModel.USER
	case model.ADMIN:
		return repoModel.ADMIN
	default:
		return repoModel.UNKNOWN
	}
}

func RepoToServiceRole(role repoModel.Role) model.Role {
	switch role {
	case repoModel.USER:
		return model.USER
	case repoModel.ADMIN:
		return model.ADMIN
	default:
		return model.UNKNOWN
	}
}

func RepoToServiceUser(user repoModel.User) model.User {
	return model.User{
		UserId:       user.UserId,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
		Role:         RepoToServiceRole(user.Role),
		UpdatedAt:    user.UpdatedAt,
		CreatedAt:    user.CreatedAt,
	}
}
