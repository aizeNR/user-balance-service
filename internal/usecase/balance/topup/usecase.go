package topup

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/service/balance"
)

type Request struct {
	UserID  uint64
	Amount  uint64
	Comment string
}

type balanceService interface {
	Add(ctx context.Context, r balance.AddRequest) error
}

type UseCase struct {
	balanceSvc balanceService
}

func New(
	balanceSvc balanceService,
) *UseCase {
	return &UseCase{
		balanceSvc: balanceSvc,
	}
}

func (u *UseCase) TopUp(ctx context.Context, r Request) error {
	if err := u.balanceSvc.Add(ctx, balance.AddRequest{
		UserID:  r.UserID,
		Amount:  r.Amount,
		Comment: r.Comment,
	}); err != nil {
		return fmt.Errorf("balanceSvc.Add: %w", err)
	}

	return nil
}
