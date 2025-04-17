package shortener

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Url   string `json:"url"`
	Hash  string `json:"hash" gorm:"uniqueIndex"`
	Stats []Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type LinkRepository interface {
	Create(link Link) (Link, error)
	GetByHash(hash string) (Link, error)
	GetByID(id uint) (Link, error)
	Update(link Link) (Link, error)
	Delete(id uint) error
}
