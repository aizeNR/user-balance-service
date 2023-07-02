package balance

import (
	"encoding/json"
	"net/http"

	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/topup"
)

type topUpRequest struct {
	UserID  uint64 `json:"user_id" validate:"required"` 
	Amount  uint64 `json:"amount" validate:"required"`
	Comment string `json:"comment" validate:"required"`
}

func (s *Server) TopUp(w http.ResponseWriter, r *http.Request) {
	var request topUpRequest

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

	if err := s.ucBalance.TopUp(ctx, topup.Request{
		UserID:  request.UserID,
		Amount:  request.Amount,
		Comment: request.Comment,
	}); err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
