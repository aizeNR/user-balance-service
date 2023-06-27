package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/transfer"
)

type balanceUseCase interface {
	TopUp(ctx context.Context, userID uint64, amount uint64) error
	WriteOff(ctx context.Context, userID uint64, amount uint64) error
	Transfer(ctx context.Context, r transfer.Request) error
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
	router.Post("/v1/balance/topup", b.TopUp)
	router.Post("/v1/balance/writeoff", b.WriteOff)
	router.Post("/v1/balance/transfer", b.Transfer)
}

func (b *BalanceServer) TopUp(w http.ResponseWriter, r *http.Request) {
	var request topUpRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httprouter.SendJsonError(w, err)
	}
	defer r.Body.Close()

	if err := b.ucBalance.TopUp(r.Context(), request.UserID, request.Amount); err != nil {
		httprouter.SendJsonError(w, err)
	}

	w.WriteHeader(http.StatusOK)
}

func (b *BalanceServer) WriteOff(w http.ResponseWriter, r *http.Request) {
	var request writeOffRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httprouter.SendJsonError(w, err)
	}
	defer r.Body.Close()

	if err := b.ucBalance.WriteOff(r.Context(), request.UserID, request.Amount); err != nil {
		httprouter.SendJsonError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (b *BalanceServer) Transfer(w http.ResponseWriter, r *http.Request) {
	var request transferRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httprouter.SendJsonError(w, err)
	}
	defer r.Body.Close()

	if err := b.ucBalance.Transfer(r.Context(), transfer.Request{
		ReceiverID: request.RceiverID,
		SenderID: request.SenderID,
		Amount: request.Amount,
	}); err != nil {
		httprouter.SendJsonError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
