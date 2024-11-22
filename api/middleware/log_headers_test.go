package middleware

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogHeadersMiddleware(t *testing.T) {
	// Capture log output
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	defer log.SetOutput(nil) // Restore log output after the test

	// Dummy handler for testing
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	logMiddleware := LogHeadersMiddleware(handler)

	t.Run("Logs headers and passes request to next handler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer token")
		rec := httptest.NewRecorder()

		logMiddleware.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert the response from the next handler
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "OK", rec.Body.String())

		// Assert the logged headers
		loggedOutput := logOutput.String()
		assert.Contains(t, loggedOutput, "Content-Type: [application/json]")
		assert.Contains(t, loggedOutput, "Authorization: [Bearer token]")
	})
}
