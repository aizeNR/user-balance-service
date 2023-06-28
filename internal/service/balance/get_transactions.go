package balance

import (
	"context"
	"fmt"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/internal/repository"
)

type GetTransactionsRequest struct {
	Paging     model.Paging
	SortFields model.TransactionSortFields
	UserID     uint64
}

type GetTransactionsResponse struct {
	Transactions []model.Transaction
}

func (s *Service) GetTransactions(ctx context.Context, r GetTransactionsRequest) (*GetTransactionsResponse, error) {
	transactions, err := s.transactionRepo.GetList(ctx, repository.GetTransactionsRequest{
		Paging:     r.Paging,
		SortFields: r.SortFields,
		UserID:     r.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("transactionRepo.GetList: %w", err)
	}

	return &GetTransactionsResponse{
		Transactions: transactions,
	}, nil
}
