package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aizeNR/user-balance-service/internal/errx"
	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/topup"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/transfer"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/writeoff"
	"github.com/go-chi/chi/v5"
)

type balanceUseCase interface {
	TopUp(ctx context.Context, r topup.Request) error
	WriteOff(ctx context.Context, r writeoff.Request) error
	Transfer(ctx context.Context, r transfer.Request) error
	GetBalance(ctx context.Context, userID uint64) (model.UserBalance, error)
}

type BalanceServer struct {
	ucBalance balanceUseCase
}

func NewBalanceServer(ucBalance balanceUseCase) *BalanceServer {
	return &BalanceServer{
		ucBalance: ucBalance,
	}
}

func (b *BalanceServer) Register(router *httprouter.Router) {
	router.Post("/v1/user/balance/topup", b.TopUp)
	router.Post("/v1/user/balance/writeoff", b.WriteOff)
	router.Post("/v1/user/balance/transfer", b.Transfer)
	router.Get("/v1/uset/{userID}/balance", b.GetBalance)
}

func (b *BalanceServer) TopUp(w http.ResponseWriter, r *http.Request) {
	var request topUpRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}
	defer r.Body.Close()

	// TODO validate.

	if err := b.ucBalance.TopUp(r.Context(), topup.Request{
		UserID:  request.UserID,
		Amount:  request.Amount,
		Comment: request.Comment,
	}); err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (b *BalanceServer) WriteOff(w http.ResponseWriter, r *http.Request) {
	var request writeOffRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}
	defer r.Body.Close()

	// TODO validate.

	if err := b.ucBalance.WriteOff(r.Context(), writeoff.Request{
		UserID:  request.UserID,
		Amount:  request.Amount,
		Comment: request.Comment,
	}); err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (b *BalanceServer) Transfer(w http.ResponseWriter, r *http.Request) {
	var request transferRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httprouter.SendJSONError(w, err)
		return
	}
	defer r.Body.Close()

	// TODO validate.

	if err := b.ucBalance.Transfer(r.Context(), transfer.Request{
		ReceiverID: request.RceiverID,
		SenderID:   request.SenderID,
		Amount:     request.Amount,
	}); err != nil {
		httprouter.SendJSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO refactoring validation.
func (b *BalanceServer) GetBalance(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		httprouter.SendJSONError(w, errx.ErrValidation{
			Fields: map[string]any{
				"userID": "empty string",
			},
		})
		return
	}

	user, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		httprouter.SendJSONError(w, errx.ErrValidation{
			Fields: map[string]any{
				"userID": "invalid string",
			},
		})
		return
	}

	balance, err := b.ucBalance.GetBalance(r.Context(), user)
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
