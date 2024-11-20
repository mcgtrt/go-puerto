package handlers

import (
	"github.com/mcgtrt/go-puerto/storage"
	"github.com/mcgtrt/go-puerto/templates/pages"
)

type ViewHandler struct {
	Store *storage.Store
}

func NewViewHandler(store *storage.Store) *ViewHandler {
	return &ViewHandler{
		Store: store,
	}
}

func (h *ViewHandler) HandleHomePage(c *Ctx) error {
	return c.Render(pages.HomePage())
}
