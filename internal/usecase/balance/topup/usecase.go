package topup

import (
	"context"
	"fmt"
)

type balanceService interface {
	Add(ctx context.Context, userID uint64, amount int64) error
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

func (u *UseCase) TopUp(ctx context.Context, userID uint64, amount int64) error {
	if err := u.balanceSvc.Add(ctx, userID, amount); err != nil {
		return fmt.Errorf("balanceSvc.Add: %w", err)
	}

	return nil
}
