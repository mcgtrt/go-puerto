package middleware

import (
	"log"
	"net/http"
)

func LogHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for name, values := range r.Header {
			log.Printf("%s: %s\n", name, values)
		}
		next.ServeHTTP(w, r)
	})
}
