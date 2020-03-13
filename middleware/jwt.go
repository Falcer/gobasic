package middleware

import "net/http"

func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// We Do some Logical/Magical here
		next.ServeHTTP(w, req)
	})
}
