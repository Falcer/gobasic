package middleware

import "net/http"

// JWT middleware
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// We Do some Logical/Magical here
		next.ServeHTTP(w, req)
	})
}
