package user

import (
	"context"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	"github.com/HeyReyHR/twitch-clone/iam/internal/repository/convert"
	passwordUtils "github.com/HeyReyHR/twitch-clone/iam/internal/utils/password"
	"github.com/HeyReyHR/twitch-clone/iam/internal/utils/validate"
)

func (s *service) Register(ctx context.Context, email, username string, role model.Role, password, passwordConfirmation string) (string, error) {
	if err := validate.RegistrationInput(email, username, password); err != nil {
		return "", err
	}

	if password != passwordConfirmation {
		return "", model.ErrPasswordMismatch
	}

	if _, err := s.repository.GetViaEmail(ctx, email); err == nil {
		return "", model.ErrEmailAlreadyTaken
	}

	if _, err := s.repository.GetViaUsername(ctx, username); err == nil {
		return "", model.ErrUsernameAlreadyTaken
	}

	hashedPassword, err := passwordUtils.HashPassword(password)
	if err != nil {
		return "", model.ErrPasswordHashFailed
	}

	userId, err := s.repository.Create(ctx, email, username, convert.ServiceToRepoRole(role), hashedPassword)
	if err != nil {
		return "", err
	}

	return userId, nil
}
