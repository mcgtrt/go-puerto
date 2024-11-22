package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeadersMiddleware(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handlerWithMiddleware := SecureHeadersMiddleware(dummyHandler)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	handlerWithMiddleware.ServeHTTP(rec, req)

	resp := rec.Result()
	defer resp.Body.Close()

	expectedHeaders := map[string]string{
		"Content-Security-Policy":   "default-src 'self'",
		"Strict-Transport-Security": "max-age=31536000; includeSubDomains",
		"X-Content-Type-Options":    "nosniff",
		"X-Frame-Options":           "DENY",
		"Referrer-Policy":           "no-referrer",
		"Permissions-Policy":        "geolocation=(), microphone=()",
	}

	for key, expectedValue := range expectedHeaders {
		value := resp.Header.Get(key)
		if value != expectedValue {
			t.Errorf("Header %q: got %q, want %q", key, value, expectedValue)
		}
	}
}
