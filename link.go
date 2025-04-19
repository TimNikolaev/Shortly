package shortener

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"gorm.io/gorm"
)

type LinkRepository interface {
	Create(link *Link) (*Link, error)
	GetByHash(hash string) (Link, error)
	GetByID(id uint) (Link, error)
	Update(link Link) (Link, error)
	Delete(id uint) error
}

type Link struct {
	gorm.Model
	URL   string `json:"url"`
	Hash  string `json:"hash" gorm:"uniqueIndex"`
	Stats []Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		URL: url,
	}
	link.GenerateHash(url)
	return link
}

func (link *Link) GenerateHash(url string) {
	link.Hash = RandStringRunes(6, url)
}

func RandStringRunes(n int, url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	hashBytes := hasher.Sum(nil)

	hashStr := base64.URLEncoding.EncodeToString(hashBytes)

	hashStr = strings.TrimRight(hashStr, "=")
	hashStr = strings.ReplaceAll(hashStr, "/", "_")

	shortHash := hashStr[:n+1]

	return shortHash
}
