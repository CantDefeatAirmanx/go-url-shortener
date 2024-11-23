package main

import (
	"log"
	"url_shortener/configs"
	"url_shortener/internal/models"
	"url_shortener/pkg/gorm_db"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := gorm_db.NewDb("links_db", &gorm_db.Config{Dsn: config.Db.Dsn})
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}

	db.AutoMigrate(&models.Link{})
}
