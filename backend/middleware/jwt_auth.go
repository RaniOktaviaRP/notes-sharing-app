package middleware

import (
	"net/http"
	"notes-app/backend/helper"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("X-API-Key")
		if tokenString == "" {
			helper.WriteUnauthorized(w, "Missing token")
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			helper.WriteUnauthorized(w, "Invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func JWTAuthHttprouter(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenString := r.Header.Get("X-API-Key")
		if tokenString == "" {
			helper.WriteUnauthorized(w, "Missing token")
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			helper.WriteUnauthorized(w, "Invalid token")
			return
		}

		h(w, r, ps)
	}
}
