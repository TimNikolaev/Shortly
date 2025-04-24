package repository

import (
	"fmt"
	"shortener"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StatPostgres struct {
	db *gorm.DB
}

func NewStatPostgres(db *gorm.DB) *StatPostgres {
	return &StatPostgres{db: db}
}

func (r *StatPostgres) AddClick(linkID uint) {
	var stat shortener.Stat
	currentDate := datatypes.Date(time.Now())
	r.db.Find(&stat, "link_id = ? AND date_stat = ?", linkID, currentDate)
	if stat.ID == 0 {
		r.db.Create(&shortener.Stat{
			LinkID: linkID,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		r.db.Save(&stat)
	}
}

func (r *StatPostgres) GetStats(linkID uint, by string, from, to time.Time) ([]shortener.GetStatResponse, error) {
	var stats []shortener.GetStatResponse
	var selectQuery string

	switch by {
	case shortener.GroupByDay:
		selectQuery = "to_char(date_stat, 'YYYY-MM-DD') as period, sum(clicks) as clicks"
	case shortener.GroupByMonth:
		selectQuery = "to_char(date_stat, 'YYYY-MM') as period, sum(clicks) as clicks"
	default:
		return nil, fmt.Errorf("invalid group by value: %v", by)
	}

	query := r.db.Table("stats").
		Select(selectQuery).
		Where("date_stat BETWEEN ? AND ?", from, to)

	if linkID == 0 {
		return nil, fmt.Errorf("invalid link_id value %v", linkID)
	}
	query = query.Where("link_id = ?", linkID)

	err := query.
		Group("period").
		Order("period").
		Scan(&stats).Error

	if err != nil {
		return nil, err
	}

	return stats, nil
}
