package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/IwanPlamboyan/contact-manajemen-golang/utils"
)

type contextKey string

const UserIDKey contextKey = "userID"

type JWTMiddleware struct {
    jwtUtil *utils.JWTUtil
}

func NewJWTMiddleware(jwtUtil *utils.JWTUtil) *JWTMiddleware {
    return &JWTMiddleware{jwtUtil: jwtUtil}
}

func (m *JWTMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := m.jwtUtil.Validate(tokenStr)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), UserIDKey, uint(userID))
        next.ServeHTTP(w, r.WithContext(ctx))
	})
}