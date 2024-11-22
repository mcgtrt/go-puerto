package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestRateLimitMiddleware(t *testing.T) {
	// Adjust rate limiter for testing
	originalLimit := DEFAULT_RATE_LIMITER_LIMIT
	originalBurst := DEFAULT_RATE_LIMITER_BURST

	DEFAULT_RATE_LIMITER_LIMIT = 2 // Limit to 2 requests per second
	DEFAULT_RATE_LIMITER_BURST = 2 // Allow up to 2 bursts
	Limiter = rate.NewLimiter(DEFAULT_RATE_LIMITER_LIMIT, DEFAULT_RATE_LIMITER_BURST)

	defer func() {
		// Restore original limiter settings after the test
		DEFAULT_RATE_LIMITER_LIMIT = originalLimit
		DEFAULT_RATE_LIMITER_BURST = originalBurst
		Limiter = rate.NewLimiter(DEFAULT_RATE_LIMITER_LIMIT, DEFAULT_RATE_LIMITER_BURST)
	}()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	rateLimitedHandler := RateLimitMiddleware(handler)

	t.Run("Within limit", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		// First request
		rateLimitedHandler.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "OK", rec.Body.String())

		// Second request
		rec = httptest.NewRecorder()
		rateLimitedHandler.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "OK", rec.Body.String())
	})

	t.Run("Exceed limit", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		// Third request exceeds the limit
		rateLimitedHandler.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusTooManyRequests, rec.Code)
		assert.Equal(t, "Too Many Requests\n", rec.Body.String())
	})

	t.Run("After reset", func(t *testing.T) {
		// Wait for limiter to reset
		time.Sleep(1 * time.Second)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		// Request after reset should succeed
		rateLimitedHandler.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "OK", rec.Body.String())
	})
}
