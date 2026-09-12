package middleware

import (
	"context"
	"net/http"
	"strings"
	"todo/pkg/jwt"
)

type key string

const (
	ContextUserId key = "userId"
)

func Unauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			Unauthed(w)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		isValid, data := jwt.NewJWT(secret).Parse(token)
		if !isValid {
			Unauthed(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserId, data.UserId)

		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
