package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mcgtrt/go-puerto/utils"
	"github.com/stretchr/testify/assert"
)

func TestLocalisationMiddleware(t *testing.T) {
	tests := []struct {
		name         string
		headers      map[string]string
		cookies      []*http.Cookie
		expectedLang string
		expectedCurr string
	}{
		{
			name:         "Default values when no headers or cookies are set",
			headers:      nil,
			cookies:      nil,
			expectedLang: DEFAULT_LANGUAGE,
			expectedCurr: DEFAULT_CURRENCY,
		},
		{
			name: "Custom Accept-Language header",
			headers: map[string]string{
				"Accept-Language": "fr",
			},
			cookies:      nil,
			expectedLang: "fr",
			expectedCurr: DEFAULT_CURRENCY,
		},
		{
			name:    "Custom currency cookie",
			headers: nil,
			cookies: []*http.Cookie{
				{Name: "currency", Value: "USD"},
			},
			expectedLang: DEFAULT_LANGUAGE,
			expectedCurr: "USD",
		},
		{
			name: "Both Accept-Language header and currency cookie provided",
			headers: map[string]string{
				"Accept-Language": "de",
			},
			cookies: []*http.Cookie{
				{Name: "currency", Value: "EUR"},
			},
			expectedLang: "de",
			expectedCurr: "EUR",
		},
		{
			name: "Missing currency cookie",
			headers: map[string]string{
				"Accept-Language": "es",
			},
			cookies: []*http.Cookie{
				{Name: "some-other-cookie", Value: "irrelevant"},
			},
			expectedLang: "es",
			expectedCurr: DEFAULT_CURRENCY,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				lang, curr := utils.GetLocale(r.Context())
				assert.Equal(t, tt.expectedLang, lang, "Expected language to match")
				assert.Equal(t, tt.expectedCurr, curr, "Expected currency to match")
			})

			middleware := LocalisationMiddleware(mockHandler)
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}
			for _, cookie := range tt.cookies {
				req.AddCookie(cookie)
			}

			rec := httptest.NewRecorder()
			middleware.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code, "Middleware should not alter the response status")
		})
	}
}
