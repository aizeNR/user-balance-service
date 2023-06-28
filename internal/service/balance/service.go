package balance

import (
	"context"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/internal/repository"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
)

type balanceRepository interface {
	Down(ctx context.Context, userID, amount uint64) error
	GetByUserID(ctx context.Context, userID uint64) (model.UserBalance, error)
	Add(ctx context.Context, userID, amount uint64) error
}

type transactionRepository interface {
	Add(ctx context.Context, transaction model.Transaction) error
	GetList(ctx context.Context, r repository.GetTransactionsRequest) ([]model.Transaction, error)
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
