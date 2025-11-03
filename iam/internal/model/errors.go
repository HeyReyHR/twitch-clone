package model

import (
	"errors"

	platformErrors "github.com/HeyReyHR/twitch-clone/platform/pkg/middleware/grpc/error"
)

// UserInput errors
var (
	ErrInvalidUsername = platformErrors.NewInvalidArgumentError(errors.New("invalid input for username"))
	ErrWeakPassword    = platformErrors.NewInvalidArgumentError(errors.New("password is too weak (minimum 8 chars, must contain a-z AND A-Z AND number)"))
	ErrInvalidEmail    = platformErrors.NewInvalidArgumentError(errors.New("this is not an email"))
)

// Register errors
var (
	ErrEmailAlreadyTaken    = platformErrors.NewInvalidArgumentError(errors.New("email is already taken"))
	ErrUsernameAlreadyTaken = platformErrors.NewInvalidArgumentError(errors.New("username is already taken"))
	ErrPasswordMismatch     = platformErrors.NewInvalidArgumentError(errors.New("password mismatch"))
)

// Auth errors
var (
	ErrInvalidCredentials = platformErrors.NewInvalidArgumentError(errors.New("invalid credentials"))
	ErrMalformedToken     = platformErrors.NewUnauthenticatedError(errors.New("token's malformed"))
)

// Server function errors
var (
	ErrGenerateTokenFailed = platformErrors.NewInternalError(errors.New("generate token failed"))
	ErrPasswordHashFailed  = platformErrors.NewInternalError(errors.New("password hash failed"))
)

// PostgreSQL errors
var (
	ErrDeletionFailed       = platformErrors.NewInternalError(errors.New("could not delete entity"))
	ErrDbScanFailed         = platformErrors.NewInternalError(errors.New("db scan failed"))
	ErrDbEntityNotFound     = platformErrors.NewNotFoundError(errors.New("db entity not found"))
	ErrCreateDbEntityFailed = platformErrors.NewInternalError(errors.New("create entity failed"))
)
