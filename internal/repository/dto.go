package repository

import "github.com/aizeNR/user-balance-service/internal/model"

type GetTransactionsRequest struct {
	Paging     model.Paging
	SortFields model.TransactionSortFields
	UserID     uint64
}
