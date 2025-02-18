package auth

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Auth())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		data := LoginResponse{
			Message: body.Email,
			Token:   handler.Config.Auth.Secret,
		}
		res.Json(w, data, 200)
		fmt.Println("Auth")
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		data := LoginResponse{
			Message: body.Email,
			Token:   handler.Config.Auth.Secret,
		}
		res.Json(w, data, 200)
		fmt.Println("Register")
	}
}
