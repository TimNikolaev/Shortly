package postgres

import (
	"shortly/internal/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(config *configs.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
