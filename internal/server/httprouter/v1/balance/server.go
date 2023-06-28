package balance

import (
	"context"
	"net/http"
	"strconv"

	"github.com/aizeNR/user-balance-service/internal/errx"
	"github.com/aizeNR/user-balance-service/internal/model"
	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
	"github.com/aizeNR/user-balance-service/internal/usecase/balance/gettransactions"
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
	GetTransactions(ctx context.Context, r gettransactions.Request) (*gettransactions.Response, error)
}

type Server struct {
	ucBalance balanceUseCase
}

func NewServer(ucBalance balanceUseCase) *Server {
	return &Server{
		ucBalance: ucBalance,
	}
}

func (s *Server) Register(router *httprouter.Router) {
	router.Post("/v1/user/balance/topup", s.TopUp)
	router.Post("/v1/user/balance/writeoff", s.WriteOff)
	router.Post("/v1/user/balance/transfer", s.Transfer)
	router.Get("/v1/user/{userID}/balance", s.GetBalance)
	router.Get("/v1/user/{userID}/balance/transactions", s.GetTransactions)
}

func (s *Server) getUserID(r *http.Request) (uint64, error) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		return 0, errx.ErrValidation{
			Fields: map[string]any{
				"userID": "empty string",
			},
		}
	}

	user, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, errx.ErrValidation{
			Fields: map[string]any{
				"userID": "invalid string",
			},
		}
	}

	return user, nil
}
