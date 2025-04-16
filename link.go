package shortener

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Url  string
	Hash string
}

type LinkRepository interface {
	Create(link Link) (Link, error)
	GetByHash(hash string) (Link, error)
	GetByID(id uint) (Link, error)
	Update(link Link) (Link, error)
	Delete(id uint) error
}
