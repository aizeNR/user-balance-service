package postgres

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
)


type UserBalanceRepository struct {
	conn postgresql.ConnManager
}

func NewUserBalanceRepository(conn postgresql.ConnManager) *UserBalanceRepository {
	return &UserBalanceRepository{
		conn: conn,
	}
}

const userBalanceTable = "user_balance"

func (u *UserBalanceRepository) Add(ctx context.Context, userID uint64, amount uint64) error {
	sql := `
	INSERT INTO user_balance 
	VALUES ($1, $2) 
	ON CONFLICT (user_id) DO UPDATE SET balance = user_balance.balance + excluded.balance
	`

	cmdTag, err := u.conn.Conn(ctx).Exec(ctx, sql, userID, amount); 
	if err != nil {
		return fmt.Errorf("failed execute query: %w", err)
	}

	// TODO CHECK THIS
	if cmdTag.RowsAffected() == 0 {
		// TODO return error
		return nil
	}

	return nil
}

func (u *UserBalanceRepository) Down(ctx context.Context, userID uint64, amount uint64) error {
	sql := `
	UPDATE user_balance SET balance = balance - $1 WHERE user_id = $2
	`

	cmdTag, err := u.conn.Conn(ctx).Exec(ctx, sql, amount, userID); 
	if err != nil {
		return fmt.Errorf("failed execute query: %w", err)
	}

	// TODO CHECK THIS
	if cmdTag.RowsAffected() == 0 {
		// TODO return error
		return nil
	}

	return nil
}

func (u *UserBalanceRepository) GetByUserID(ctx context.Context, userID uint64) (model.UserBalance, error) {
	sql := `
	SELECT u.user_id, u.balance 
	FROM user_balance as u
		WHERE user_id = $1
	FOR UPDATE
	`

	var balance model.UserBalance

	if err := u.conn.Conn(ctx).QueryRow(ctx, sql, userID).Scan(&balance.UserID, &balance.Balance); err != nil {
		return model.UserBalance{}, fmt.Errorf("failed execute query: %w", err)
	}

	return balance, nil
}
