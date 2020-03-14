package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arganaphangquestian/gobasic/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

// JWT middleware
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Header["token"] != nil {
			token, err := jwt.Parse(req.Header["token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error")
				}
				return []byte("argadev123"), nil
			})
			if err != nil {
				json.NewEncoder(w).Encode(utils.ErrorResponse("Token invalid", nil))
			}
			if token.Valid {
				next.ServeHTTP(w, req)
			}
		} else {
			json.NewEncoder(w).Encode(utils.ErrorResponse("Token invalid", nil))
		}
	})
}
