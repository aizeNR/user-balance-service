package sort

import (
	"net/http"

	"github.com/aizeNR/user-balance-service/internal/model"
)

type sortTransactionField interface {
	GetKey() string
	Apply(value string, tsf *model.TransactionSortFields)
}

type operationDate struct{}

func (operationDate) GetKey() string {
	return "sort[operation_date]"
}

func (operationDate) Apply(value string, tsf *model.TransactionSortFields) {
	d, ok := model.SortMap[value]
	if !ok {
		return
	}

	tsf.OperationDate = d
}

type amountSortField struct{}

func (amountSortField) GetKey() string {
	return "sort[amount]"
}

func (amountSortField) Apply(value string, tsf *model.TransactionSortFields) {
	d, ok := model.SortMap[value]
	if !ok {
		return
	}

	tsf.OperationDate = d
}

var transactionFields = []sortTransactionField{
	amountSortField{},
	operationDate{},
}

func ExtractTransactionFields(r *http.Request) model.TransactionSortFields {
	values := r.URL.Query()
	if len(values) == 0 {
		return model.TransactionSortFields{}
	}

	sortFields := &model.TransactionSortFields{}
	for _, v := range transactionFields {
		value := r.URL.Query().Get(v.GetKey())
		if value == "" {
			continue
		}

		v.Apply(value, sortFields)
	}

	return *sortFields
}
