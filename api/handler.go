package api

import "github.com/mcgtrt/go-puerto/storage"

type Handler struct {
	Store *storage.Store
}

func NewHandler(store *storage.Store) *Handler {
	return &Handler{
		Store: store,
	}
}
