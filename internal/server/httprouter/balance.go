package httprouter

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type balanceUseCase interface {
	TopUp(ctx context.Context, userID uint64, amount int64) error
	WriteOff(ctx context.Context, userID uint64, amount int64) error
}

type BalanceServer struct {
	ucBalance balanceUseCase
}

func NewBalanceServer(ucBalance balanceUseCase) *BalanceServer {
	return &BalanceServer{
		ucBalance: ucBalance,
	}
}

func (b *BalanceServer) Register(router chi.Router) {
	router.Post("/v1/balance/topup", b.TopUp)
}

func (b *BalanceServer) TopUp(w http.ResponseWriter, r *http.Request) {
	
}
