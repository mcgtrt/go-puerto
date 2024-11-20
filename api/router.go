package api

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()

	mountMiddlewares(r)
	mountRoutes(r, h)

	return r
}

func mountMiddlewares(r *chi.Mux) {}

func mountRoutes(r *chi.Mux, h *Handler) {}
