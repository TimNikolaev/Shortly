package shortener

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Url  string
	Hash string
}
