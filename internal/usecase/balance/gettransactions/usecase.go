package gettransactions

import (
	"context"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/internal/service/balance"
)

type Request struct {
	// TODO change to cursor paginate.
	Paging     model.Paging
	SortFields model.TransactionSortFields
	UserID     uint64
}

type Response struct {
	Transactions []model.Transaction
}

type balanceService interface {
	GetTransactions(ctx context.Context, r balance.GetTransactionsRequest) (*balance.GetTransactionsResponse, error)
}

type UseCase struct {
	balanceSvc balanceService
}

func New(
	balanceSvc balanceService,
) *UseCase {
	return &UseCase{
		balanceSvc: balanceSvc,
	}
}

func (u *UseCase) GetTransactions(ctx context.Context, r Request) (*Response, error) {
	resp, err := u.balanceSvc.GetTransactions(ctx, balance.GetTransactionsRequest{
		Paging:     r.Paging,
		SortFields: r.SortFields,
		UserID:     r.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		Transactions: resp.Transactions,
	}, nil
}
