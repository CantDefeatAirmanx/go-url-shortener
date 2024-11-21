package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

type Config struct {
	Dsn string
}

func NewDb(conf *Config) (*Db, error) {
	db, err := gorm.Open(postgres.Open(conf.Dsn), &gorm.Config{})

	return &Db{db}, err
}
