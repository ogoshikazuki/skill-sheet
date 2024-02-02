package entity

import (
	"context"
	"log"

	"github.com/cockroachdb/errors"
)

var (
	ErrInternal = errors.New("internal server error")
)

func ErrInternalAndLogStack(_ context.Context, err error) error {
	log.Printf("%+v", errors.WithStack(err))

	return ErrInternal
}
