package balance

import (
	"encoding/json"
	"net/http"

	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/transfer"
)

type transferRequest struct {
	RceiverID uint64 `json:"receiver_id"`
	SenderID  uint64 `json:"sender_id"`
	Amount    uint64 `json:"amount"`
	Comment   string `json:"comment"`
}

func (s *Server) Transfer(w http.ResponseWriter, r *http.Request) {
	var request transferRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}
	defer r.Body.Close()

	// TODO validate.

	if err := s.ucBalance.Transfer(r.Context(), transfer.Request{
		ReceiverID: request.RceiverID,
		SenderID:   request.SenderID,
		Amount:     request.Amount,
	}); err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
