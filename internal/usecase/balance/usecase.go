package balance

import (
	"context"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/gettransactions"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/topup"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/transfer"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/writeoff"
)

type topUpAction interface {
	TopUp(ctx context.Context, r topup.Request) error
}

type writeOffAction interface {
	WriteOff(ctx context.Context, r writeoff.Request) error
}

type transferAction interface {
	Transfer(ctx context.Context, r transfer.Request) error
}

type getBalanceAction interface {
	GetBalance(ctx context.Context, userID uint64) (model.UserBalance, error)
}

type getTransactionsAction interface {
	GetTransactions(ctx context.Context, r gettransactions.Request) (*gettransactions.Response, error)
}

type UseCase struct {
	topUpAction
	writeOffAction
	transferAction
	getBalanceAction
	getTransactionsAction
}

func New(
	topUp topUpAction,
	writeOff writeOffAction,
	transfer transferAction,
	getBalance getBalanceAction,
	getTransactions getTransactionsAction,
) *UseCase {
	return &UseCase{
		topUpAction:           topUp,
		writeOffAction:        writeOff,
		transferAction:        transfer,
		getBalanceAction:      getBalance,
		getTransactionsAction: getTransactions,
	}
}
