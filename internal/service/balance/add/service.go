package add

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/pkg/clock"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
	"github.com/gofrs/uuid/v5"
)

type balanceRepository interface {
	Add(ctx context.Context, userID, amount uint64) error
}

type transactionRepository interface {
	Add(ctx context.Context, transaction model.Transaction) error
}

type Service struct {
	balanceRepo     balanceRepository
	transactionRepo transactionRepository
	txManager       postgresql.TransactionManager
}

func NewService(
	balanceRepo balanceRepository,
	transactionRepo transactionRepository,
	txManager postgresql.TransactionManager,
) *Service {
	return &Service{
		balanceRepo:     balanceRepo,
		transactionRepo: transactionRepo,
		txManager:       txManager,
	}
}

func (u *Service) Add(ctx context.Context, userID, amount uint64) error {
	return u.txManager.RunTx(ctx, func(ctx context.Context) error {
		if err := u.balanceRepo.Add(ctx, userID, amount); err != nil {
			return fmt.Errorf("balanceRepo.Add: %w", err)
		}

		// TODO interface to generator
		transactionID, err := uuid.NewV4()
		if err != nil {
			return err
		}

		err = u.transactionRepo.Add(ctx, model.Transaction{
			ID:            transactionID,
			UserID:        userID,
			Amount:        int64(amount),
			OperationDate: clock.Now(),
		})
		if err != nil {
			return fmt.Errorf("transactionRepo.Add: %w", err)
		}

		return nil
	})
}
