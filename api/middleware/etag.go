package middleware

import "net/http"

func ETagMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eTag := `W/"123456"`
		if match := r.Header.Get("If-None-Match"); match == eTag {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.Header().Set("ETag", eTag)
		next.ServeHTTP(w, r)
	})
}
