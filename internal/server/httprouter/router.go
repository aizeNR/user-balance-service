package httprouter

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Init() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	return r
}
