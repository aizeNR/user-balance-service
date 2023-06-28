package service

import (
	"github.com/aizeNR/user-balance-service/internal/repository/postgres"
	"github.com/aizeNR/user-balance-service/internal/service/balance"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
)

type BalanceDeps struct {
	TxManager *postgresql.Manager
}

func InitBalance(deps *BalanceDeps) *balance.Service {
	balanceRepo := postgres.NewUserBalanceRepository(deps.TxManager)
	transactionRepo := postgres.NewTransactionRepository(deps.TxManager)

	return balance.NewService(
		balanceRepo,
		transactionRepo,
		deps.TxManager,
	)
}
