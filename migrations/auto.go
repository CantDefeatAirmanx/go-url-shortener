package main

import (
	"log"
	"url_shortener/configs"
	"url_shortener/internal/models/link"
	"url_shortener/pkg/gorm"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := gorm.NewDb(&gorm.Config{Dsn: config.Db.Dsn})
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}

	db.AutoMigrate(&link.Link{})
}
