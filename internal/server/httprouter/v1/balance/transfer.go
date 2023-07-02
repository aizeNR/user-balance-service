package balance

import (
	"encoding/json"
	"net/http"

	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/transfer"
)

type transferRequest struct {
	RceiverID uint64 `json:"receiver_id" validate:"required"`
	SenderID  uint64 `json:"sender_id" validate:"required"`
	Amount    uint64 `json:"amount" validate:"required"`
	Comment   string `json:"comment" validate:"required"`
}

func (s *Server) Transfer(w http.ResponseWriter, r *http.Request) {
	var request transferRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	ctx := r.Context()

	err = s.validator.StructCtx(ctx, request)
	if err != nil {
		httprouter.SendValidationError(w, err)
		return
	}

	if err := s.ucBalance.Transfer(ctx, transfer.Request{
		ReceiverID: request.RceiverID,
		SenderID:   request.SenderID,
		Amount:     request.Amount,
	}); err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
