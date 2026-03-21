package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Siwani-tech/GoAuth-Lite.git/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const Useremailkey contextKey = "useremail"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		ctx := context.WithValue(r.Context(), Useremailkey, email)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
