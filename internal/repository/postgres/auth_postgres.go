package postgres

import (
	"errors"
	"shortly"

	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user *shortly.User) (uint, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r *UserPostgres) GetUser(email, password_hash string) (*shortly.User, error) {
	var user shortly.User

	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	if !checkPasswordHash(password_hash, user.Password) {
		return nil, errors.New("error wrong password")
	}

	return &user, nil
}

func checkPasswordHash(password_hash, userPassword_hash string) bool {
	return userPassword_hash == password_hash
}
