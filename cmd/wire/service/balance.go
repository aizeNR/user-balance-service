package service

import (
	"github.com/aizeNR/user-balance-service/internal/repository/postgres"
	"github.com/aizeNR/user-balance-service/internal/service/balance"
	"github.com/aizeNR/user-balance-service/internal/service/balance/add"
	"github.com/aizeNR/user-balance-service/internal/service/balance/down"
	"github.com/aizeNR/user-balance-service/internal/service/balance/getbyuserid"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
)

type BalanceDeps struct {
	TxManager *postgresql.Manager
}

func InitBalance(deps *BalanceDeps) *balance.Service {
	balanceRepo := postgres.NewUserBalanceRepository(deps.TxManager)
	transactionRepo := postgres.NewTransactionRepository(deps.TxManager)

	addAction := add.NewService(balanceRepo, transactionRepo, deps.TxManager)
	downAction := down.NewService(balanceRepo, transactionRepo, deps.TxManager)
	getByUserIDAction := getbyuserid.NewService(balanceRepo)

	return balance.New(
		addAction,
		downAction,
		getByUserIDAction,
	)
}
