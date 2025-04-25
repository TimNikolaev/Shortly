package shortly

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatRepository interface {
	AddClick(linkId uint)
	GetStats(linkID uint, by string, from, to time.Time) ([]GetStatResponse, error)
}

type Stat struct {
	gorm.Model
	LinkID uint           `json:"link_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Clicks int            `json:"clicks"`
	Date   datatypes.Date `json:"date" gorm:"column:date_stat"`
}

type GetStatResponse struct {
	Period string `json:"period"`
	Clicks int    `json:"sum"`
}
