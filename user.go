package shortener

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	CreateUser(user User) (int, error)
	GetUser(email, password string) (User, error)
}
