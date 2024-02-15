package repository

import (
	"context"

	"github.com/ogoshikazuki/skill-sheet/entity"
)

type transactionController struct {
	sqlHandler SqlHandler
}

func NewTransactionController(sqlHandler SqlHandler) entity.TransactionController {
	return &transactionController{
		sqlHandler: sqlHandler,
	}
}

func (t *transactionController) Transaction(ctx context.Context, f func(tx entity.Tx) error) error {
	tx, err := t.sqlHandler.BeginTx(ctx)
	if err != nil {
		return err
	}

	if err := f(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
