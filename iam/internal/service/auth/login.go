package auth

import (
	"context"
	"errors"

	"github.com/HeyReyHR/twitch-clone/iam/internal/model"
	"github.com/HeyReyHR/twitch-clone/iam/internal/repository/convert"
	repoModel "github.com/HeyReyHR/twitch-clone/iam/internal/repository/model"
	"github.com/HeyReyHR/twitch-clone/iam/internal/utils/jwt_tokens"
	passwordUtils "github.com/HeyReyHR/twitch-clone/iam/internal/utils/password"
	"github.com/HeyReyHR/twitch-clone/iam/internal/utils/validate"
)

func (s *service) Login(ctx context.Context, login, password string) (*model.TokenPair, error) {
	var user repoModel.User
	var tokens *model.TokenPair
	err := validate.LoginInput(login, password)
	if err != nil {
		if errors.Is(err, model.ErrInvalidCredentials) {
			return nil, err
		}
		if errors.Is(err, model.ErrInvalidEmail) {
			user, err = s.userRepository.GetViaUsername(ctx, login)
			if err != nil {
				if errors.Is(err, model.ErrDbEntityNotFound) {
					return nil, model.ErrInvalidCredentials
				}
				return nil, model.ErrDbScanFailed
			}
		}
	} else {
		user, err = s.userRepository.GetViaEmail(ctx, login)
		if err != nil {
			if errors.Is(err, model.ErrDbEntityNotFound) {
				return nil, model.ErrInvalidCredentials
			}
			return nil, model.ErrDbScanFailed
		}
	}

	isPassword := passwordUtils.CheckPasswordHash(password, user.PasswordHash)
	if !isPassword {
		return nil, model.ErrInvalidCredentials
	}

	tokens, err = jwt_tokens.GenerateTokenPair(convert.RepoToServiceUser(user), s.accessTokenTtl, s.refreshTokenTtl)
	if err != nil {
		return nil, err
	}

	if _, err = s.authRepository.CreateRefreshToken(ctx, user.UserId, tokens.RefreshToken, s.refreshTokenTtl); err != nil {
		return nil, err
	}

	return tokens, err
}
