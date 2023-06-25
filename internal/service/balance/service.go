package balance

import "context"


type addAction interface {
	Add(ctx context.Context, userID uint64, amount int64) error
}

type downAction interface {
	Down(ctx context.Context, userID uint64, amount int64) error
}

type Service struct {
	addAction
	downAction
}

func New(add addAction, down downAction) *Service {
	return &Service{
		addAction: add,
		downAction: down,
	}
}
