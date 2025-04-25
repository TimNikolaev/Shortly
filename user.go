package shortly

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required" gorm:"column:password_hash"`
	Links    []Link `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserRepository interface {
	CreateUser(user *User) (uint, error)
	GetUser(email, password string) (*User, error)
}
