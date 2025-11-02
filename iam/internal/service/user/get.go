package user

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	"github.com/HeyReyHR/twitch-clone/iam/internal/repository/convert"
)

func (s *service) Get(ctx context.Context, userId string) (model.User, error) {
	user, err := s.repository.Get(ctx, userId)
	if err != nil {
		return model.User{}, err
	}
	
	return convert.RepoToServiceUser(user), nil
}
