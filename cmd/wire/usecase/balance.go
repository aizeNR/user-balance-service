package usecase

import (
	balanceSvc "github.com/aizeNR/user-balance-service/internal/service/balance"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/topup"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/transfer"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/writeoff"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
)

type BalanceDeps struct {
	BalanceService *balanceSvc.Service
	TxManager *postgresql.Manager
}

func InitBalance(deps *BalanceDeps) *balance.UseCase {
	return balance.New(
		topup.New(deps.BalanceService),
		writeoff.New(deps.BalanceService),
		transfer.New(deps.BalanceService, deps.TxManager),
	)
}
