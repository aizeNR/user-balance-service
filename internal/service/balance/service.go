package balance

import (
	"context"

	"github.com/aizeNR/user-balance-service/internal/model"
)

type addAction interface {
	Add(ctx context.Context, userID, amount uint64) error
}

type downAction interface {
	Down(ctx context.Context, userID, amount uint64) error
}

type getByUserIDAction interface {
	GetByUserID(ctx context.Context, userID uint64) (model.UserBalance, error)
}

type Service struct {
	addAction
	downAction
	getByUserIDAction
}

func New(add addAction, down downAction, get getByUserIDAction) *Service {
	return &Service{
		addAction:         add,
		downAction:        down,
		getByUserIDAction: get,
	}
}
