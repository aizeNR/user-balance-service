package balance

import "context"

type topUpAction interface {
	TopUp(ctx context.Context, userID uint64, amount int64) error
}

type writeOffAction interface {
	WriteOff(ctx context.Context, userID uint64, amount int64) error
}

type UseCase struct {
	topUpAction
	writeOffAction
}

func New(
	t topUpAction,
	w writeOffAction,
) *UseCase {
	return &UseCase{
		topUpAction: t,
		writeOffAction: w,
	}
}
