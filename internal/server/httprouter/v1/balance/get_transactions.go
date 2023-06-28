package balance

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
	"github.com/aizeNR/user-balance-service/internal/server/httprouter/v1/balance/sort"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/gettransactions"
	"github.com/gofrs/uuid/v5"
)

type getTransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID            uuid.UUID `json:"id"`
	UserID        uint64    `json:"user_id"`
	Amount        int64     `json:"amount"`
	OperationDate time.Time `json:"operation_date"`
	Comment       string    `json:"comment"`
}

func (s *Server) GetTransactions(w http.ResponseWriter, r *http.Request) {
	userID, err := s.getUserID(r)
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	res, err := s.ucBalance.GetTransactions(r.Context(), gettransactions.Request{
		Paging: model.Paging{
			Limit: httprouter.ExtractLimit(r),
			Page:  httprouter.ExtractPage(r),
		},
		UserID:     userID,
		SortFields: sort.ExtractTransactionFields(r),
	})
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	// TODO mapper
	resp := make([]Transaction, 0, len(res.Transactions))
	for _, v := range res.Transactions {
		resp = append(resp, Transaction{
			ID:            v.ID,
			UserID:        v.UserID,
			Amount:        v.Amount,
			OperationDate: v.OperationDate,
			Comment:       v.Comment,
		})
	}

	answer, err := json.Marshal(getTransactionsResponse{
		Transactions: resp,
	})
	if err != nil {
		httprouter.SendJSONError(w, err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}
