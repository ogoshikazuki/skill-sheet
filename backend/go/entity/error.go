package entity

import "github.com/cockroachdb/errors"

type InternalServerError struct {
	errWithStack error
}

var _ error = (*InternalServerError)(nil)

func (e *InternalServerError) Error() string {
	return "internal server error"
}

func (e *InternalServerError) ErrWithStack() error {
	return e.errWithStack
}

func NewInternalServerError(err error) *InternalServerError {
	return &InternalServerError{
		errWithStack: errors.WithStack(err),
	}
}
