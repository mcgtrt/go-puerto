package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRateLimitMiddleware(t *testing.T) {
	// Configure rate limiter settings for the test
	originalLimit := DEFAULT_RATE_LIMITER_LIMIT
	originalBurst := DEFAULT_RATE_LIMITER_BURST

	// Adjust limiter for testing
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

	// Helper function to send a request
	sendRequest := func() *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		rateLimitedHandler.ServeHTTP(rec, req)
		return rec
	}

	// Test within the limit
	for i := 0; i < 2; i++ { // Send two requests (limit is 2 per second)
		resp := sendRequest()
		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code 200, got %d", resp.Code)
		}
	}

	// Test exceeding the limit
	resp := sendRequest() // Third request should fail
	if resp.Code != http.StatusTooManyRequests {
		t.Errorf("Expected status code 429, got %d", resp.Code)
	}

	// Wait for the rate limiter to reset
	time.Sleep(time.Second)

	// Test after waiting for the limiter to reset
	resp = sendRequest()
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d after reset", resp.Code)
	}
}
