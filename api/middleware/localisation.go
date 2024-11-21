package middleware

import (
	"context"
	"net/http"

	"github.com/mcgtrt/go-puerto/types"
)

const (
	DEFAULT_LANGUAGE = "en"
	DEFAULT_CURRENCY = "GBP"
)

func LocalisationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Accept-Language")
		if lang == "" {
			lang = DEFAULT_LANGUAGE
		}

		currency, err := r.Cookie("currency")
		if err != nil {
			// fallback
			currency = &http.Cookie{
				Name:  "currency",
				Value: DEFAULT_CURRENCY,
			}
		}

		ctx := context.WithValue(r.Context(), types.LanguageCtxKey{}, lang)
		ctx = context.WithValue(ctx, types.CurrencyCtxKey{}, currency.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
