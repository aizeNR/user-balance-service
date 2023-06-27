package balance

import (
	"context"

	"github.com/aizeNR/user-balance-service/internal/usecase/balance/transfer"
)

type topUpAction interface {
	TopUp(ctx context.Context, userID uint64, amount uint64) error
}

type writeOffAction interface {
	WriteOff(ctx context.Context, userID uint64, amount uint64) error
}

type transferAction interface {
	Transfer(ctx context.Context, r transfer.Request) error
}

type UseCase struct {
	topUpAction
	writeOffAction
	transferAction
}

func New(
	topUp topUpAction,
	writeOff writeOffAction,
	transfer transferAction,
) *UseCase {
	return &UseCase{
		topUpAction: topUp,
		writeOffAction: writeOff,
		transferAction: transfer,
	}
}
