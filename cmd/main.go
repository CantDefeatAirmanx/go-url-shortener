package main

import (
	"fmt"
	"log"
	"net/http"
	"url_shortener/configs"
	"url_shortener/internal/api/auth/auth_v1"
	links_v1 "url_shortener/internal/api/links/v1"
	"url_shortener/pkg/gorm"
)

func main() {
	router := http.NewServeMux()

	config := configs.GetConfig()

	auth_v1.NewAuthHandlerV1(router, auth_v1.AuthHandlerConfig(config.Auth))
	links_v1.NewLinksHandlerV1(router)

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
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	_, err = gorm.NewDb(&gorm.Config{Dsn: config.Db.Dsn})
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}
}
