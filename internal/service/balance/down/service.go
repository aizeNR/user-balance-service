package down

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/errx"
	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/pkg/clock"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
	"github.com/gofrs/uuid/v5"
)

type balanceRepository interface {
	Down(ctx context.Context, userID, amount uint64) error
	GetByUserID(ctx context.Context, userID uint64) (model.UserBalance, error)
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

func (u *Service) Down(ctx context.Context, userID, amount uint64) error {
	return u.txManager.RunTx(ctx, func(ctx context.Context) error {
		balance, err := u.balanceRepo.GetByUserID(ctx, userID)
		if err != nil {
			return fmt.Errorf("balanceRepo.GetByUserID: %w", err)
		}

		if int64(balance.Balance)-int64(amount) < 0 {
			return &errx.ErrNotEnoughtMoney{}
		}

		if err := u.balanceRepo.Down(ctx, userID, amount); err != nil {
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
			Amount:        (-1 * int64(amount)),
			OperationDate: clock.Now(),
		})
		if err != nil {
			return fmt.Errorf("transactionRepo.Add: %w", err)
		}

		return nil
	})
}
