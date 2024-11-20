package middleware

import (
	"context"
	"net/http"
)

const (
	DEFAULT_LANGUAGE = "en"
	DEFAULT_CURRENCY = "GBP"
)

type LocaleCtx struct{}
type CurrencyCtx struct{}

func LocaleMiddleware(next http.Handler) http.Handler {
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

		ctx := context.WithValue(r.Context(), LocaleCtx{}, lang)
		ctx = context.WithValue(ctx, CurrencyCtx{}, currency.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
