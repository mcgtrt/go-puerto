package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

var DEFAULT_RATE_LIMITER_LIMIT rate.Limit = 5
var DEFAULT_RATE_LIMITER_BURST = 15

var Limiter = rate.NewLimiter(DEFAULT_RATE_LIMITER_LIMIT, DEFAULT_RATE_LIMITER_BURST)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !Limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
