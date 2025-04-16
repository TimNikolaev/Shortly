package repository

import (
	"shortener"

	"gorm.io/gorm"
)

type LinkPostgres struct {
	db *gorm.DB
}

func NewLinkPostgres(db *gorm.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (s *LinkPostgres) Create(link shortener.Link) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *LinkPostgres) GetByHash(hash string) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *LinkPostgres) GetByID(id uint) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *LinkPostgres) Update(link shortener.Link) (shortener.Link, error) {
	return shortener.Link{}, nil
}

func (s *LinkPostgres) Delete(id uint) error {
	return nil
}
