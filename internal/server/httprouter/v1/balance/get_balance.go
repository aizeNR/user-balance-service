package balance

import (
	"encoding/json"
	"net/http"

	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
)

type getBalanceResponse struct {
	UserID  uint64 `json:"user_id"`
	Balance uint64 `json:"balance"`
}

func (s *Server) GetBalance(w http.ResponseWriter, r *http.Request) {
	userID, err := s.getUserID(r)
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	balance, err := s.ucBalance.GetBalance(r.Context(), userID)
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	resp := getBalanceResponse{
		UserID:  balance.UserID,
		Balance: balance.Balance,
	}

	answer, err := json.Marshal(resp)
	if err != nil {
		httprouter.SendJSONError(w, err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}
