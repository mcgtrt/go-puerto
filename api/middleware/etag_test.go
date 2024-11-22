package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestETagMiddleware(t *testing.T) {
	// A dummy handler to test the middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})
	eTagHandler := ETagMiddleware(handler)

	t.Run("Without If-None-Match header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		eTagHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert that the ETag header is set
		assert.Equal(t, `W/"123456"`, resp.Header.Get("ETag"))
		// Assert the status code and response body
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	})

	t.Run("With matching If-None-Match header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("If-None-Match", `W/"123456"`) // Set a matching ETag
		rec := httptest.NewRecorder()

		eTagHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert status code for Not Modified
		assert.Equal(t, http.StatusNotModified, resp.StatusCode)
		// Assert that no body is returned
		assert.Empty(t, rec.Body.String())
	})

	t.Run("With non-matching If-None-Match header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("If-None-Match", `W/"654321"`) // Set a non-matching ETag
		rec := httptest.NewRecorder()

		eTagHandler.ServeHTTP(rec, req)

		resp := rec.Result()
		defer resp.Body.Close()

		// Assert that the ETag header is set
		assert.Equal(t, `W/"123456"`, resp.Header.Get("ETag"))
		// Assert the status code and response body
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	})
}
