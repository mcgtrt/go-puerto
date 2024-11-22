package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateHeadersMiddleware(t *testing.T) {
	// Dummy handler to test the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	validateHeadersHandler := ValidateHeadersMiddleware(handler)

	t.Run("Valid Content-Type header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json") // Set valid header
		rec := httptest.NewRecorder()

		validateHeadersHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert response status and body
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "OK", rec.Body.String())
	})

	t.Run("Invalid Content-Type header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "text/plain") // Set invalid header
		rec := httptest.NewRecorder()

		validateHeadersHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert response status and error message
		assert.Equal(t, http.StatusUnsupportedMediaType, resp.StatusCode)
		assert.Equal(t, "Invalid Content-Type\n", rec.Body.String())
	})

	t.Run("Missing Content-Type header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil) // No Content-Type header
		rec := httptest.NewRecorder()

		validateHeadersHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert response status and error message
		assert.Equal(t, http.StatusUnsupportedMediaType, resp.StatusCode)
		assert.Equal(t, "Invalid Content-Type\n", rec.Body.String())
	})
}
