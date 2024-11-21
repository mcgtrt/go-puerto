package api

import (
	"github.com/mcgtrt/go-puerto/api/handlers"
	"github.com/mcgtrt/go-puerto/storage"
)

type Handler struct {
	View *handlers.ViewHandler
}

func NewHandler(store *storage.Store) *Handler {
	return &Handler{
		View: handlers.NewViewHandler(store),
	}
}
