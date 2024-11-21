package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/mcgtrt/go-puerto/api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestMountFileServer(t *testing.T) {
	// Create a temporary static directory for testing
	wd, _ := os.Getwd()
	staticPath := filepath.Join(wd, "static_test")
	err := os.MkdirAll(staticPath, 0755) // Use MkdirAll to avoid "file exists" errors
	assert.NoError(t, err, "Failed to create static directory")
	defer os.RemoveAll(staticPath) // Cleanup after the test

	// Create a mock file in the static directory
	filePath := filepath.Join(staticPath, "test.txt")
	err = os.WriteFile(filePath, []byte("test content"), 0644)
	assert.NoError(t, err, "Failed to create test file")

	// Initialize the router and mount the file server
	r := chi.NewRouter()
	mountFileServer(r, "static_test", "static_test")

	// Test serving an existing file
	req := httptest.NewRequest(http.MethodGet, "/static_test/test.txt", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Assert 200 OK and correct file content
	assert.Equal(t, http.StatusOK, w.Code, "Expected 200 OK response for static file")
	assert.Equal(t, "test content", w.Body.String(), "Expected file content to be served")

	// Test serving a non-existent file
	req = httptest.NewRequest(http.MethodGet, "/static_test/nonexistent.txt", nil)
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Assert 404 Not Found
	assert.Equal(t, http.StatusNotFound, w.Code, "Expected 404 Not Found for non-existent file")
}

func TestWrap(t *testing.T) {
	fn := func(c *handlers.Ctx) error {
		c.Text(http.StatusOK, "success")
		return nil
	}

	// Test successful execution
	t.Run("Success case", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		handler := wrap(fn)
		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected 200 OK response")
		assert.Equal(t, "success", w.Body.String(), "Expected response body to match")
	})

	// Test error handling
	t.Run("Error case", func(t *testing.T) {
		fn := func(c *handlers.Ctx) error {
			c.Error(http.StatusInternalServerError)
			return nil
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		handler := wrap(fn)
		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected 500 Internal Server Error response")
	})
}
