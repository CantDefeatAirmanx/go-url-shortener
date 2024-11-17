package auth

import (
	"fmt"
	"net/http"
)

type AuthHandler struct{}

const prefix = "/api/v1/auth"

func NewAuthHandler(router *http.ServeMux) {
	handler := AuthHandler{}

	router.HandleFunc(fmt.Sprintf("POST %s/login", prefix), handler.handleLogin())
	router.HandleFunc(fmt.Sprintf("POST %s/register", prefix), handler.handleRegister())
}

func (handler *AuthHandler) handleLogin() http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		fmt.Println("/login")
		resWriter.Write([]byte("works"))
	}
}

func (handler *AuthHandler) handleRegister() http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		fmt.Println("/register")
	}
}
