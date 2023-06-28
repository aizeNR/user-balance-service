package balance

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/errx"
	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/pkg/clock"
	"github.com/gofrs/uuid/v5"
)

type DownRequest struct {
	UserID  uint64
	Amount  uint64
	Comment string
}

func (s *Service) Down(ctx context.Context, r DownRequest) error {
	return s.txManager.RunTx(ctx, func(ctx context.Context) error {
		balance, err := s.balanceRepo.GetByUserID(ctx, r.UserID)
		if err != nil {
			return fmt.Errorf("balanceRepo.GetByUserID: %w", err)
		}

		if int64(balance.Balance)-int64(r.Amount) < 0 {
			return &errx.ErrNotEnoughtMoney{}
		}

		if err := s.balanceRepo.Down(ctx, r.UserID, r.Amount); err != nil {
			return fmt.Errorf("balanceRepo.Add: %w", err)
		}

		// TODO interface to id generator
		transactionID, err := uuid.NewV4()
		if err != nil {
			return err
		}

		err = s.transactionRepo.Add(ctx, model.Transaction{
			ID:            transactionID,
			UserID:        r.UserID,
			Amount:        (-1 * int64(r.Amount)),
			OperationDate: clock.Now(),
		})
		if err != nil {
			return fmt.Errorf("transactionRepo.Add: %w", err)
		}

		return nil
	})
}
