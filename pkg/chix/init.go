package chix

import (
	"github.com/go-chi/chi/v5"
)

func NewChiMux() *chi.Mux {
	r := chi.NewRouter()
	return r
}
