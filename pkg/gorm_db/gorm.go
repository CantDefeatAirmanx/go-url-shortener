package gorm_db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

type Config struct {
	Dsn string
}

var dbMap map[string]*Db = make(map[string]*Db)

func NewDb(key string, conf *Config) (*Db, error) {
	if _, ok := dbMap[key]; ok {
		panic("db with key " + key + " already inited")
	}

	db, err := gorm.Open(postgres.Open(conf.Dsn), &gorm.Config{})
	dbInstance := Db{db}

	dbMap[key] = &dbInstance

	return &dbInstance, err
}

func GetDb(key string) (*Db, error) {
	db, ok := dbMap[key]

	if !ok {
		return nil, fmt.Errorf("there is no db by key %s", key)
	}

	return db, nil
}
