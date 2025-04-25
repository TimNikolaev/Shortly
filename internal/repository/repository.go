package repository

import (
	"shortly"

	"gorm.io/gorm"
)

type Repository struct {
	shortly.LinkRepository
	shortly.UserRepository
	shortly.StatRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		LinkRepository: NewLinkPostgres(db),
		UserRepository: NewUserPostgres(db),
		StatRepository: NewStatPostgres(db),
	}
}
