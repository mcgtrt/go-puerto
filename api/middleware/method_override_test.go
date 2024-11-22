package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodOverrideMiddleware(t *testing.T) {
	// Dummy handler for testing
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(r.Method)) // Respond with the method to test the override
	})

	methodOverrideHandler := MethodOverrideMiddleware(handler)

	t.Run("Without X-HTTP-Method-Override header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		methodOverrideHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert the response method matches the original method
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, http.MethodGet, rec.Body.String())
	})

	t.Run("With X-HTTP-Method-Override header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("X-HTTP-Method-Override", http.MethodPut) // Override POST with PUT
		rec := httptest.NewRecorder()

		methodOverrideHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert the response method matches the overridden method
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, http.MethodPut, rec.Body.String())
	})

	t.Run("With invalid method in X-HTTP-Method-Override", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("X-HTTP-Method-Override", "INVALID") // Set an invalid method
		rec := httptest.NewRecorder()

		methodOverrideHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert the response method matches the invalid override (middleware does not validate methods)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "INVALID", rec.Body.String())
	})
}
