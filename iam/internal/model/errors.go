package model

import (
	"errors"

	platformErrors "github.com/HeyReyHR/twitch-clone/platform/pkg/middleware/grpc/error"
)

var (
	ErrUserScanFailed       = platformErrors.NewInternalError(errors.New("user scan failed"))
	ErrCreateUserFailed     = platformErrors.NewInternalError(errors.New("create user failed"))
	ErrPasswordHashFailed   = platformErrors.NewInternalError(errors.New("password hash failed"))
	ErrUserNotFound         = platformErrors.NewNotFoundError(errors.New("user not found"))
	ErrInvalidUsername      = platformErrors.NewInvalidArgumentError(errors.New("invalid input for username"))
	ErrWeakPassword         = platformErrors.NewInvalidArgumentError(errors.New("password is too weak (minimum 8 chars, must contain a-z AND A-Z AND number)"))
	ErrInvalidEmail         = platformErrors.NewInvalidArgumentError(errors.New("this is not an email"))
	ErrEmailAlreadyTaken    = platformErrors.NewInvalidArgumentError(errors.New("email is already taken"))
	ErrUsernameAlreadyTaken = platformErrors.NewInvalidArgumentError(errors.New("username is already taken"))
	ErrPasswordMismatch     = platformErrors.NewInvalidArgumentError(errors.New("password mismatcg"))
	ErrInvalidCredentials   = platformErrors.NewInvalidArgumentError(errors.New("invalid credentials"))
)
