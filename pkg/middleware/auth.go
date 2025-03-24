package middleware

import (
	"context"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)
		// fmt.Println("TOKEN -", token)
		// fmt.Println("ISVALID -", isValid)
		// fmt.Println("DATA -", data)
		next.ServeHTTP(w, req)
	})
}
