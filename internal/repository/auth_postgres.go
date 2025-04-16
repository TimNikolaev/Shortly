package repository

import (
	"shortener"

	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (s *UserPostgres) CreateUser(user shortener.User) (int, error) {
	return 0, nil
}

func (s *UserPostgres) GetUser(email, password string) (shortener.User, error) {
	return shortener.User{}, nil
}
