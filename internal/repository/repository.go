package repository

import (
	"shortly"
	"shortly/internal/repository/postgres"

	"gorm.io/gorm"
)

type Repository struct {
	shortly.LinkRepository
	shortly.UserRepository
	shortly.StatRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		LinkRepository: postgres.NewLinkPostgres(db),
		UserRepository: postgres.NewUserPostgres(db),
		StatRepository: postgres.NewStatPostgres(db),
	}
}
