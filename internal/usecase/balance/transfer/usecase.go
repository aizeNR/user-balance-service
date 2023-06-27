package transfer

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/pkg/postgresql"
)

type balanceService interface {
	Add(ctx context.Context, userID uint64, amount uint64) error
	Down(ctx context.Context, userID uint64, amount uint64) error
}

type Request struct {
	ReceiverID uint64
	SenderID uint64
	Amount uint64
}

type UseCase struct {
	balanceSvc balanceService
	txManager postgresql.TransactionManager
}

func New(
	balanceSvc balanceService,
	txManager postgresql.TransactionManager,
) *UseCase {
	return &UseCase{
		balanceSvc: balanceSvc,
		txManager: txManager,
	}
}

func (u *UseCase) Transfer(ctx context.Context, r Request) error {
	return u.txManager.RunTx(ctx, func(ctx context.Context) error {
		if err := u.balanceSvc.Down(ctx, r.SenderID, r.Amount); err != nil {
			return fmt.Errorf("balanceSvc.Down: %w", err)
		}

		if err := u.balanceSvc.Add(ctx, r.ReceiverID, r.Amount); err != nil {
			return fmt.Errorf("balanceSvc.Add: %w", err)
		}

		return nil
	})
}
