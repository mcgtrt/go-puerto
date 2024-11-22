package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORSMiddleware(t *testing.T) {
	// A dummy handler to test the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	corsHandler := CORSMiddleware(handler)

	t.Run("Regular request adds CORS headers", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		corsHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert CORS headers
		assert.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "GET, POST, PUT, DELETE, OPTIONS", resp.Header.Get("Access-Control-Allow-Methods"))
		assert.Equal(t, "Content-Type, Authorization", resp.Header.Get("Access-Control-Allow-Headers"))
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Assert handler response
		assert.Equal(t, "OK", rec.Body.String())
	})

	t.Run("Preflight OPTIONS request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodOptions, "/", nil)
		rec := httptest.NewRecorder()

		corsHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert CORS headers
		assert.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Origin"))
		assert.Equal(t, "GET, POST, PUT, DELETE, OPTIONS", resp.Header.Get("Access-Control-Allow-Methods"))
		assert.Equal(t, "Content-Type, Authorization", resp.Header.Get("Access-Control-Allow-Headers"))

		// Assert preflight response
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
		assert.Empty(t, rec.Body.String()) // OPTIONS requests have no body
	})
}
