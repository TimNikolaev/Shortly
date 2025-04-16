package repository

import (
	"shortener/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewPostgresDB(config *configs.Config) (*DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}
