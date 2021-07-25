package chix

import (
	"github.com/go-chi/chi/v5"
)

func NewChiMux() *chi.Mux {
	r := chi.NewRouter()
	//r.Use(chiMiddleware.RequestID)
	//r.Use(middleware.NewStructuredLogger(logger))
	//r.Use(chiMiddleware.Recoverer)

	return r
}
