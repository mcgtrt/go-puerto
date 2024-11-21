package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/mcgtrt/go-puerto/api/handlers"
	"github.com/mcgtrt/go-puerto/api/middleware"
	"github.com/mcgtrt/go-puerto/utils"
)

// API func signature for the handler methods
type APIFunc func(c *handlers.Ctx) error

// Returns a fully mounted Chi router
func NewRouter(h *Handler, cfg *utils.Config) *chi.Mux {
	r := chi.NewRouter()

	mountMiddlewares(r, cfg.Middleware)
	mountRoutes(r, h, cfg.HTTP)

	return r
}

// The place to mount all the middlewares
func mountMiddlewares(r *chi.Mux, cfg *utils.MiddlewareConfig) {
	if cfg.Localisation {
		r.Use(middleware.LocalisationMiddleware)
	}
	// TODO: For the scope of the next release
	// if cfg.SecureHeaders {
	// }
	// if cfg.RateLimit {
	// }
	// if cfg.LogAndMonitorHeaders {
	// }
	// if cfg.CORS {
	// }
	// if cfg.ETAG {
	// }
	// if cfg.ValidateSanitiseHeaders {
	// }
	// if cfg.MethodOverride {
	// }
}

// This is the global routes mount entry. Add new mountSomethig
// into this method to keep it simple and nicely organised
func mountRoutes(r *chi.Mux, h *Handler, cfg *utils.HTTPConfig) {
	if cfg.FileServerPath != "" {
		mountFileServer(r, cfg.FileServerPath, "static")
	}
	mountView(r, h.View)
}

// Ensure local file server is serving the files from the directory
// that will have the same url path in accessing the file server
func mountFileServer(r *chi.Mux, pathURL, staticDir string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	static := filepath.Join(wd, staticDir)
	fs := http.FileServer(http.Dir(static))
	fileServer := http.StripPrefix("/"+pathURL, fs)

	r.Get("/"+pathURL+"/*", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})
}

// Use to match all the routes and implement serving web pages
func mountView(r *chi.Mux, h *handlers.ViewHandler) {
	r.Get("/", wrap(h.HandleHomePage))
}

// Use this function to convert APIFunc to http.HandlerFunc
// and handle possible errors with the Error Handler func
func wrap(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := handlers.NewCtx(w, r)

		if err := fn(ctx); err != nil {
			// TODO: HANDLE ERRORS
			ctx.Error(http.StatusInternalServerError)
		}
	}
}
