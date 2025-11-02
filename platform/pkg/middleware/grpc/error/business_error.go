package error

import (
	"github.com/go-faster/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorCode int64

const (
	NotFoundErrorCode ErrorCode = iota
	InvalidArgumentErrorCode
	UnauthenticatedErrorCode
	InternalErrorCode
)

type businessError struct {
	code ErrorCode
	err  error
}

func (b *businessError) Error() string {
	if b.err != nil {
		return b.err.Error()
	}
	return "unknown business error"
}

func (b *businessError) Unwrap() error {
	return b.err
}

func (b *businessError) Code() ErrorCode {
	return b.code
}

func NewNotFoundError(err error) *businessError {
	return &businessError{
		code: NotFoundErrorCode,
		err:  err,
	}
}

func NewInvalidArgumentError(err error) *businessError {
	return &businessError{
		code: InvalidArgumentErrorCode,
		err:  err,
	}
}

func NewUnauthenticatedError(err error) *businessError {
	return &businessError{
		code: UnauthenticatedErrorCode,
		err:  err,
	}
}

func NewInternalError(err error) *businessError {
	return &businessError{
		code: InternalErrorCode,
		err:  err,
	}
}

func GetBusinessError(err error) *businessError {
	var businessErr *businessError
	if errors.As(err, &businessErr) {
		return businessErr
	}
	return nil
}

func errorCodeToGRPCCode(code ErrorCode) codes.Code {
	switch code {
	case UnauthenticatedErrorCode:
		return codes.Unauthenticated
	case InvalidArgumentErrorCode:
		return codes.InvalidArgument
	case NotFoundErrorCode:
		return codes.NotFound
	case InternalErrorCode:
		return codes.Internal
	default:
		return codes.Unknown
	}
}

func BusinessErrorToGRPCStatus(err *businessError) *status.Status {
	grpcCode := errorCodeToGRPCCode(err.Code())
	return status.New(grpcCode, err.Error())
}
