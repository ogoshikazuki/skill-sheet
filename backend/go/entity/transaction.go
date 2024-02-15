package entity

//go:generate go run github.com/matryer/moq -out transaction_controller_mock.go . TransactionController Tx

import "context"

type TransactionController interface {
	Transaction(ctx context.Context, f func(tx Tx) error) error
}

type Tx interface{}
