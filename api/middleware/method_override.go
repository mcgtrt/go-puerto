package middleware

import "net/http"

func MethodOverrideMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if override := r.Header.Get("X-HTTP-Method-Override"); override != "" {
			r.Method = override
		}
		next.ServeHTTP(w, r)
	})
}
