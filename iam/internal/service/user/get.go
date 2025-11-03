package user

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	"github.com/HeyReyHR/twitch-clone/iam/internal/repository/convert"
	"github.com/HeyReyHR/twitch-clone/iam/internal/utils"
)

func (s *service) Get(ctx context.Context, userId string) (*model.User, error) {
	user, err := s.repository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	return utils.Pointer[model.User](convert.RepoToServiceUser(user)), nil
}
