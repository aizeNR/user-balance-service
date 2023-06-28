package transfer

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/service/balance"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
)

type balanceService interface {
	Add(ctx context.Context, r balance.AddRequest) error
	Down(ctx context.Context, r balance.DownRequest) error
}

type Request struct {
	ReceiverID uint64
	SenderID   uint64
	Amount     uint64
	Comment    string
}

type UseCase struct {
	balanceSvc balanceService
	txManager  postgresql.TransactionManager
}

func New(
	balanceSvc balanceService,
	txManager postgresql.TransactionManager,
) *UseCase {
	return &UseCase{
		balanceSvc: balanceSvc,
		txManager:  txManager,
	}
}

func (u *UseCase) Transfer(ctx context.Context, r Request) error {
	return u.txManager.RunTx(ctx, func(ctx context.Context) error {
		err := u.balanceSvc.Down(ctx, balance.DownRequest{
			UserID:  r.SenderID,
			Amount:  r.Amount,
			Comment: r.Comment,
		})
		if err != nil {
			return fmt.Errorf("balanceSvc.Down: %w", err)
		}

		err = u.balanceSvc.Add(ctx, balance.AddRequest{
			UserID:  r.ReceiverID,
			Amount:  r.Amount,
			Comment: r.Comment,
		})
		if err != nil {
			return fmt.Errorf("balanceSvc.Add: %w", err)
		}

		return nil
	})
}
