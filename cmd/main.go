package main

import (
	"fmt"
	"net/http"
	"url_shortener/configs"
	"url_shortener/internal/api/auth/v1"
	"url_shortener/pkg/db"
)

func main() {
	router := http.NewServeMux()

	config := configs.GetConfig()

	auth.NewAuthHandler(router, auth.AuthHandlerConfig(config.Auth))

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
	config := configs.LoadConfig()

	_, dbErr := db.NewDb(&db.Config{Dsn: config.Db.Dsn})
	if dbErr != nil {
		panic(dbErr)
	}
}
