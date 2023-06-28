package postgres

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
)

type TransactionRepository struct {
	conn postgresql.ConnManager
}

func NewTransactionRepository(conn postgresql.ConnManager) *TransactionRepository {
	return &TransactionRepository{
		conn: conn,
	}
}

const transactionTable = "transactions"

func (t *TransactionRepository) Add(ctx context.Context, transaction model.Transaction) error {
	sql, args, err := postgresql.Builder.Insert(transactionTable).
		Columns(
			"id",
			"user_id",
			"amount",
			"operation_date",
			"comment",
		).Values(
		transaction.ID,
		transaction.UserID,
		transaction.Amount,
		transaction.OperationDate,
		transaction.Comment,
	).ToSql()
	if err != nil {
		return fmt.Errorf("failed execute query: %w", err)
	}

	_, err = t.conn.Conn(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed execute query: %w", err)
	}

	return nil
}
