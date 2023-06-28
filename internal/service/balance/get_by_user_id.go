package balance

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/model"
)

func (s *Service) GetByUserID(ctx context.Context, userID uint64) (model.UserBalance, error) {
	balance, err := s.balanceRepo.GetByUserID(ctx, userID)
	if err != nil {
		return model.UserBalance{}, fmt.Errorf("balanceRepo.GetByUserID: %w", err)
	}

	return balance, nil
}
