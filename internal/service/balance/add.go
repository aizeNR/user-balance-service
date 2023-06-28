package balance

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/pkg/clock"
	"github.com/gofrs/uuid/v5"
)

type AddRequest struct {
	UserID  uint64
	Amount  uint64
	Comment string
}

func (s *Service) Add(ctx context.Context, r AddRequest) error {
	return s.txManager.RunTx(ctx, func(ctx context.Context) error {
		if err := s.balanceRepo.Add(ctx, r.UserID, r.Amount); err != nil {
			return fmt.Errorf("balanceRepo.Add: %w", err)
		}

		// TODO interface to generator
		transactionID, err := uuid.NewV4()
		if err != nil {
			return err
		}

		err = s.transactionRepo.Add(ctx, model.Transaction{
			ID:            transactionID,
			UserID:        r.UserID,
			Amount:        int64(r.Amount),
			OperationDate: clock.Now(),
			Comment:       r.Comment,
		})
		if err != nil {
			return fmt.Errorf("transactionRepo.Add: %w", err)
		}

		return nil
	})
}
