package middleware

import (
	"net/http"
	"strings"
	"os"
	"notes-app/backend/helper"

	"github.com/golang-jwt/jwt/v5"
)

func getSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	return r.Header.Get("X-API-Key")
}

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		if tokenString == "" {
			helper.WriteUnauthorized(w, "Missing token")
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return getSecret(), nil
		})

		if err != nil || !token.Valid {
			helper.WriteUnauthorized(w, "Invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
