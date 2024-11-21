package handlers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCtx(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	ctx := NewCtx(rec, req)

	assert.Equal(t, req.Context(), ctx.Context, "Expected context to match request context")
	assert.Equal(t, rec, ctx.Response, "Expected response writer to match recorder")
	assert.Equal(t, req, ctx.Request, "Expected request to match original request")
}

// Mock for a a-h/templ component
type MockTemplComponent struct {
	mock.Mock
}

// Render mocks the Render method of the templ.Component interface
func (m *MockTemplComponent) Render(ctx context.Context, w io.Writer) error {
	args := m.Called(ctx, w)
	return args.Error(0)
}

func TestRender(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := NewCtx(rec, req)
	mockComponent := &MockTemplComponent{}
	mockComponent.On("Render", req.Context(), rec).Return(nil).Once()

	err := ctx.Render(mockComponent)
	assert.NoError(t, err, "Expected no error when rendering component")
	mockComponent.AssertExpectations(t)
}

func TestJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := NewCtx(rec, req)

	// Test with valid JSON response
	data := map[string]string{"key": "value"}
	err := ctx.JSON(http.StatusOK, data)

	assert.NoError(t, err, "Expected no error when encoding JSON")
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"), "Expected content type to be application/json")
	assert.Equal(t, http.StatusOK, rec.Code, "Expected status code to match")
	assert.JSONEq(t, `{"key": "value"}`, rec.Body.String(), "Expected JSON response to match")

	// Test with invalid JSON (e.g., circular reference)
	err = ctx.JSON(http.StatusOK, make(chan int)) // Invalid JSON
	assert.Error(t, err, "Expected error when encoding invalid JSON")
}

func TestText(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := NewCtx(rec, req)

	message := "Hello, World!"
	err := ctx.Text(http.StatusOK, message)

	assert.NoError(t, err, "Expected no error when writing text response")
	assert.Equal(t, "text/plain; charset=utf-8", rec.Header().Get("Content-Type"), "Expected content type to be text/plain")
	assert.Equal(t, http.StatusOK, rec.Code, "Expected status code to match")
	assert.Equal(t, message, rec.Body.String(), "Expected response body to match text")
}

func TestError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := NewCtx(rec, req)

	ctx.Error(http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected status code to match")
	assert.Equal(t, "Bad Request\n", rec.Body.String(), "Expected error message to match")
}

// MockReadCloser is a mock implementation of io.ReadCloser
type MockReadCloser struct {
	mock.Mock
}

// Read mocks the Read method of io.ReadCloser
func (m *MockReadCloser) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

// Close mocks the Close method of io.ReadCloser
func (m *MockReadCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestCloseBody(t *testing.T) {
	body := &MockReadCloser{}
	body.On("Close").Return(nil).Once()

	req := httptest.NewRequest(http.MethodPost, "/", body)
	rec := httptest.NewRecorder()
	ctx := NewCtx(rec, req)

	ctx.CloseBody()
	body.AssertExpectations(t)
}
