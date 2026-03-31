package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	CtxKeyUserID  contextKey = "userId"
	CtxKeyUsername contextKey = "username"
)

func Auth(next http.Handler) http.Handler {
	whitelist := []string{"/health", "/api/v1/user/login", "/api/v1/user/register"}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, p := range whitelist {
			if r.URL.Path == p {
				next.ServeHTTP(w, r)
				return
			}
		}

		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    401,
				"message": "unauthorized",
			})
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    401,
				"message": "invalid token",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx := r.Context()
			ctx = context.WithValue(ctx, CtxKeyUserID, claims["user_id"])
			ctx = context.WithValue(ctx, CtxKeyUsername, claims["username"])
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
