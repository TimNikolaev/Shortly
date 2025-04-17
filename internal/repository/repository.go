package repository

import (
	"shortener"

	"gorm.io/gorm"
)

type Repository struct {
	shortener.LinkRepository
	shortener.UserRepository
	shortener.StatRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		LinkRepository: NewLinkPostgres(db),
		UserRepository: NewUserPostgres(db),
		StatRepository: NewStatPostgres(db),
	}
}
