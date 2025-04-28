package shortly

import (
	"math/rand"

	"gorm.io/gorm"
)

type LinkRepository interface {
	Create(link *Link) (*Link, error)
	GetByHash(hash string) (*Link, error)
	GetByID(id uint) (*Link, error)
	GetAll(userID uint, limit, offset int) ([]Link, error)
	Count(userID uint) (int64, error)
	Update(link *Link, userID uint) (*Link, error)
	Delete(userID, linkID uint) error
}

type Link struct {
	gorm.Model
	UserID uint   `json:"link_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	URL    string `json:"url"`
	Hash   string `json:"hash" gorm:"uniqueIndex"`
	Stats  []Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(id uint, url string) *Link {
	link := &Link{
		UserID: id,
		URL:    url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = generateRandomKey(6)
}

func generateRandomKey(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	key := make([]byte, length)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}
	return string(key)
}
