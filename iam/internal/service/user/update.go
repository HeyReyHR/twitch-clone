package user

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/HeyReyHR/twitch-clone/iam/internal/config"
	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
)

func (s *service) Update(ctx context.Context, userId string, username, email *string, avatar []byte, contentType *string, isStreaming *bool) error {
	var avatarUrl *string

	if avatar != nil && contentType != nil {
		filename := fmt.Sprintf("%s_%d", userId, time.Now().Unix())
		err := s.minio.UploadFile(ctx, config.AppConfig().Minio.AvatarBucket(), filename, bytes.NewReader(avatar), int64(len(avatar)), *contentType)
		if err != nil {
			return model.ErrUploadFile
		}

		url := s.minio.GetFileUrl(config.AppConfig().Minio.AvatarBucket(), filename)
		avatarUrl = &url
	}

	err := s.repository.Update(ctx, userId, repoModel.UpdateParams{
		Username:    username,
		Email:       email,
		AvatarUrl:   avatarUrl,
		IsStreaming: isStreaming,
	})
	if err != nil {
		return err
	}

	return nil
}
