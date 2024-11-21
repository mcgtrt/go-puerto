package api

import (
	"github.com/mcgtrt/go-puerto/api/handlers"
	"github.com/mcgtrt/go-puerto/storage"
	"github.com/mcgtrt/go-puerto/utils"
)

type Handler struct {
	Config *utils.HTTPConfig
	View   *handlers.ViewHandler
}

func NewHandler(store *storage.Store, config *utils.HTTPConfig) *Handler {
	return &Handler{
		Config: config,
		View:   handlers.NewViewHandler(store),
	}
}
