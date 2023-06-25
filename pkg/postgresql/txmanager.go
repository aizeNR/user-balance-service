package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type TransactionManager interface {
	RunTx(ctx context.Context, do func(ctx context.Context) error) error
}

type ConnManager interface {
	Conn(ctx context.Context) Connection
	RunTx(ctx context.Context, do func(ctx context.Context) error) error
}

type Connection interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

func NewTxManager(client *Postgres) *Manager {
	return &Manager{
		client: client,
	}
}

type Manager struct {
	client *Postgres
}

func (s *Manager) RunTx(ctx context.Context, do func(ctx context.Context) error) error {
	_, ok := hasTx(ctx)
	if ok {
		return do(ctx)
	}

	return runTx(ctx, s.client, do)
}

func (s *Manager) Conn(ctx context.Context) Connection {
	tx, ok := hasTx(ctx)
	if ok {
		return tx.conn
	}

	return s.client.Pool
}

type txKey int

const (
	key txKey = iota
)

type transaction struct {
	conn pgx.Tx
}

func (t *transaction) commit(ctx context.Context) error {
	if err := t.conn.Commit(ctx); err != nil {
		return fmt.Errorf("transaction commit: %w", err)
	}

	return nil
}

func (t *transaction) rollback(ctx context.Context) error {
	if err := t.conn.Rollback(ctx); err != nil {
		return fmt.Errorf("transaction rollback: %w", err)
	}

	return nil
}

func withTx(ctx context.Context, tx transaction) context.Context {
	return context.WithValue(ctx, key, tx)
}

func hasTx(ctx context.Context) (transaction, bool) {
	tx, ok := ctx.Value(key).(transaction)
	if ok {
		return tx, true
	}

	return transaction{}, false
}

func runTx(ctx context.Context, client *Postgres, do func(ctx context.Context) error) error {
	conn, err := client.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction runTx: %w", err)
	}

	tx := transaction{conn: conn}

	txCtx := withTx(ctx, tx)

	err = do(txCtx)
	if err != nil {
		if err := tx.rollback(txCtx); err != nil {
			return err
		}

		return err
	}

	return tx.commit(txCtx)
}
