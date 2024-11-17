package main

import (
	"fmt"
	"net/http"
	"url_shortener/configs"
	"url_shortener/internal/api/v1/auth"
)

func main() {
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	config := configs.GetConfig()
	port := config.APP_PORT

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	fmt.Printf("Server is listening on port: %s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func init() {
	configs.LoadConfig()
}
