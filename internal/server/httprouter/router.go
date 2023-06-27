package httprouter

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const jsonType = "application/json"

type Router struct {
	*chi.Mux
}

func Init() *Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// TODO refactor
	r.Use(middleware.SetHeader("Content-Type", jsonType))

	return &Router{
		r,
	}
}
