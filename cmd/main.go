package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/aizeNR/user-balance-service/cmd/wire/service"
	"github.com/aizeNR/user-balance-service/cmd/wire/usecase"
	"github.com/aizeNR/user-balance-service/internal/config"
	"github.com/aizeNR/user-balance-service/internal/server/httprouter"
	"github.com/aizeNR/user-balance-service/internal/server/httprouter/v1/balance"
	"github.com/aizeNR/user-balance-service/pkg/httpserver"
	"github.com/aizeNR/user-balance-service/pkg/postgresql"
	"github.com/rs/zerolog"
)

func main() {
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGINT}

	doneCh := make(chan os.Signal, len(signals))

	signal.Notify(doneCh, signals...)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := zerolog.New(zerolog.NewConsoleWriter())

	cfg, err := config.New()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed init config")
	}

	// Init Database
	db, err := postgresql.New(ctx, cfg.Postgres.URL)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed init db")
	}
	defer db.Close()

	txManager := postgresql.NewTxManager(db)

	// Usecase
	balanceService := service.InitBalance(&service.BalanceDeps{
		TxManager: txManager,
	})

	ucBalance := usecase.InitBalance(&usecase.BalanceDeps{
		BalanceService: balanceService,
		TxManager:      txManager,
	})

	router := httprouter.Init()

	balanceServer := balance.NewServer(ucBalance)

	balanceServer.Register(router)

	httpServer := httpserver.New(router, httpserver.Port(cfg.HTTP.GetPort()))

	defer func() {
		err = httpServer.Shutdown()
		if err != nil {
			logger.Error().Err(err).Msg("failed http server shutdown")
		}
	}()

	select {
	case <-doneCh:
		logger.Info().Msg("gracefully stopped")
	case err := <-httpServer.Notify():
		logger.Error().Err(err).Msg("httpServer.Notify")
	}
}
