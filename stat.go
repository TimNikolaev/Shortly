package shortener

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	LinkID uint           `json:"link_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Clicks int            `json:"clicks"`
	Date   datatypes.Date `json:"date"`
}

type StatGetResponse struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}

type StatRepository interface {
	AddClick(linkId uint)
	GetStats(by string, from, to time.Time) []StatGetResponse
}
