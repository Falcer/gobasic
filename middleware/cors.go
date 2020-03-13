package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")  // For Development we allow to open api to everyone
		w.Header().Set("Access-Control-Allow-Methods", "*") // For Development we allow to open all methods functionality to everyone
		next.ServeHTTP(w, req)
	})
}
