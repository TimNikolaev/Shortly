package repository

import (
	"shortener"
	"time"

	"gorm.io/gorm"
)

type StatPostgres struct {
	db *gorm.DB
}

func NewStatPostgres(db *gorm.DB) *StatPostgres {
	return &StatPostgres{db: db}
}

func (r *StatPostgres) AddClick(linkId uint) {

}

func (r *StatPostgres) GetStats(by string, from, to time.Time) []shortener.StatGetResponse {
	return nil
}
