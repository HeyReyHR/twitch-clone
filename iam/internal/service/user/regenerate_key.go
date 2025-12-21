package user

import (
	"context"

	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
	"github.com/google/uuid"
)

func (s *service) RegenerateKey(ctx context.Context, userId string) (string, error) {
	streamKey := uuid.NewString()
	err := s.repository.Update(ctx, userId, repoModel.UpdateParams{
		StreamKey: &streamKey,
	})
	if err != nil {
		return "", err
	}

	return streamKey, nil
}
