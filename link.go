package shortener

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"gorm.io/gorm"
)

type LinkRepository interface {
	Create(link *Link) (*Link, error)
	GetByHash(hash string) (*Link, error)
	GetByID(id uint) (*Link, error)
	GetAll(userID uint, limit, offset int) ([]Link, error)
	Count(userID uint) (int64, error)
	Update(link Link) (Link, error)
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
	link.GenerateHash(url)
	return link
}

func (link *Link) GenerateHash(url string) {
	link.Hash = Hashing(6, url)
}

func Hashing(n int, url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	hashBytes := hasher.Sum(nil)

	hashStr := base64.URLEncoding.EncodeToString(hashBytes)

	hashStr = strings.TrimRight(hashStr, "=")
	hashStr = strings.ReplaceAll(hashStr, "/", "_")

	shortHash := hashStr[:n]

	return shortHash
}
