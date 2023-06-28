package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/internal/repository"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
	"github.com/jackc/pgx/v5"
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
		return fmt.Errorf("failed build query: %w", err)
	}

	_, err = t.conn.Conn(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed execute query: %w", err)
	}

	return nil
}

func (t *TransactionRepository) GetList(ctx context.Context, r repository.GetTransactionsRequest) ([]model.Transaction, error) {
	builder := postgresql.Builder.Select(
		"id",
		"user_id",
		"amount",
		"operation_date",
		"comment",
	).
		From(transactionTable).
		Where("user_id = ?", r.UserID).
		Limit(r.Paging.Limit).
		Offset(r.Paging.GetOffset())

	sql, args, err := t.mapSortFieldsToBuidler(builder, r.SortFields).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed build query: %w", err)
	}

	rows, err := t.conn.Conn(ctx).Query(ctx, sql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed execute query: %w", err)
	}
	defer rows.Close()

	transactions := make([]model.Transaction, 0, r.Paging.Limit)

	for rows.Next() {
		var transaction model.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Amount,
			&transaction.OperationDate,
			&transaction.Comment,
		)
		if err != nil {
			return nil, fmt.Errorf("failed scan: %w", err)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (t *TransactionRepository) mapSortFieldsToBuidler(b squirrel.SelectBuilder, tsf model.TransactionSortFields) squirrel.SelectBuilder {
	if tsf.Amount != "" {
		b = b.OrderBy(fmt.Sprintf("amount %s", tsf.Amount))
	}

	if tsf.OperationDate != "" {
		b = b.OrderBy(fmt.Sprintf("operation_date %s", tsf.OperationDate))
	}

	return b
}
