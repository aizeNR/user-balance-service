package writeoff

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
	Down(ctx context.Context, r balance.DownRequest) error
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

func (u *UseCase) WriteOff(ctx context.Context, r Request) error {
	if err := u.balanceSvc.Down(ctx, balance.DownRequest{
		UserID:  r.UserID,
		Amount:  r.Amount,
		Comment: r.Comment,
	}); err != nil {
		return fmt.Errorf("balanceSvc.Add: %w", err)
	}

	return nil
}
