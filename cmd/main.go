package main

import (
	"fmt"
	"log"
	"net/http"
	"url_shortener/configs"
	auth_handler_v1 "url_shortener/internal/api/auth/v1"
	links_handler_v1 "url_shortener/internal/api/links/v1"
	link_repository_v1 "url_shortener/internal/repository/link/v1"
	link_service_v1 "url_shortener/internal/service/link/v1"
	"url_shortener/pkg/gorm_db"
)

func main() {
	router := http.NewServeMux()
	config := configs.GetConfig()
	linksDb, _ := gorm_db.GetDb(LINKS_DB_KEY)

	auth_handler_v1.NewAuthHandlerV1(router, auth_handler_v1.AuthHandlerConfig(config.Auth))

	linksRepo := link_repository_v1.NewLinkRepositoryGorm(linksDb)
	linksService := link_service_v1.NewLinkService(linksRepo)
	links_handler_v1.NewLinksHandlerV1(links_handler_v1.LinksHandlerV1Deps{Router: router, Service: linksService})

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

const (
	LINKS_DB_KEY string = "links_db"
)

func init() {

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	_, err = gorm_db.NewDb(LINKS_DB_KEY, &gorm_db.Config{Dsn: config.Db.Dsn})
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}
}
