package getbyuserid

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/model"
)

type balanceRepository interface {
	GetByUserID(ctx context.Context, userID uint64) (model.UserBalance, error)
}

type Service struct {
	balanceRepo balanceRepository
}

func NewService(
	balanceRepo balanceRepository,
) *Service {
	return &Service{
		balanceRepo: balanceRepo,
	}
}

func (u *Service) GetByUserID(ctx context.Context, userID uint64) (model.UserBalance, error) {
	balance, err := u.balanceRepo.GetByUserID(ctx, userID)
	if err != nil {
		return model.UserBalance{}, fmt.Errorf("balanceRepo.GetByUserID: %w", err)
	}

	return balance, nil
}
