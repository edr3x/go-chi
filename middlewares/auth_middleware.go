package middlewares

import (
	"context"
	"fmt"
	"net/http"
)

type key string

const (
	UserDataKey key = "userData"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("RequireAuth Middleware Triggered")

		// passing the value in context
		ctx := r.Context()

		userData := "some user data"

		cx := context.WithValue(ctx, UserDataKey, userData)

		// Pass the new context to the next handler
		next.ServeHTTP(w, r.WithContext(cx))
	})
}
