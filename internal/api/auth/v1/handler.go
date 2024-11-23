package auth_handler_v1

import (
	"fmt"
	"net/http"
	"url_shortener/pkg/utils/http_helpers"
)

type AuthHandler struct {
	*AuthHandlerConfig
}

type AuthHandlerConfig struct {
	Secret string
}

const prefix = "/api/auth/v1"

func NewAuthHandlerV1(router *http.ServeMux, config AuthHandlerConfig) {
	handler := AuthHandler{&config}

	router.HandleFunc(fmt.Sprintf("POST %s/login", prefix), handler.handleLogin())
	router.HandleFunc(fmt.Sprintf("POST %s/register", prefix), handler.handleRegister())
}

func (handler *AuthHandler) handleLogin() http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		reqPayload, err := http_helpers.HandleJsonReqBody[LoginRequest](&resWriter, req)
		if err != nil {
			return
		}
		_ = reqPayload

		resPayload := LoginResponse{
			Token: "123",
		}

		http_helpers.WriteJsonRes(resWriter, resPayload, http.StatusOK)
	}
}

func (handler *AuthHandler) handleRegister() http.HandlerFunc {
	return func(resWriter http.ResponseWriter, req *http.Request) {
		reqPayload, err := http_helpers.HandleJsonReqBody[RegisterRequest](&resWriter, req)
		if err != nil {
			return
		}
		_ = reqPayload
		http_helpers.WriteJsonRes(resWriter, "ok", http.StatusOK)
	}
}
