package getbalance

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/model"
)

type balanceService interface {
	GetByUserID(ctx context.Context, userID uint64) (model.UserBalance, error)
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

func (u *UseCase) GetBalance(ctx context.Context, userID uint64) (model.UserBalance, error) {
	balance, err := u.balanceSvc.GetByUserID(ctx, userID)
	if err != nil {
		return model.UserBalance{}, fmt.Errorf("balanceSvc.GetByUserID: %w", err)
	}

	return balance, nil
}
