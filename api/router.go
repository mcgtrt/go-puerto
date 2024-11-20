package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcgtrt/go-puerto/api/handlers"
	"github.com/mcgtrt/go-puerto/api/middleware"
)

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()

	mountMiddlewares(r)
	mountRoutes(r, h)

	return r
}

func mountMiddlewares(r *chi.Mux) {
	r.Use(middleware.LocaleMiddleware)
}

func mountRoutes(r *chi.Mux, h *Handler) {
	mountFileServer(r)
	mountView(r, h.View)
}

func mountFileServer(r *chi.Mux) {

}

func mountView(r *chi.Mux, h *handlers.ViewHandler) {
	r.Get("/", handle(h.HandleHomePage))
}

// Use this function to convert APIFunc to http.HandlerFunc including Error Handling
func handle(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := handlers.NewCtx(w, r)

		if err := fn(ctx); err != nil {
			// TODO: HANDLE ERRORS
			ctx.Error(http.StatusInternalServerError)
		}
	}
}
