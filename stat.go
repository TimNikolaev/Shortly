package shortener

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type Stat struct {
	gorm.Model
	LinkID uint           `json:"link_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Clicks int            `json:"clicks"`
	Date   datatypes.Date `json:"date" gorm:"column:date_stat"`
}

type StatGetResponse struct {
	Period string `json:"period" gorm:"column:period"`
	Clicks int    `json:"sum" gorm:"column:clicks"`
}

type StatRepository interface {
	AddClick(linkId uint)
	GetStats(by string, from, to time.Time) ([]StatGetResponse, error)
}
