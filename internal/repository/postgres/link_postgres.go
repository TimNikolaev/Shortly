package postgres

import (
	"fmt"
	"shortly"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LinkPostgres struct {
	db *gorm.DB
}

func NewLinkPostgres(db *gorm.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (r *LinkPostgres) Create(link *shortly.Link) (*shortly.Link, error) {
	result := r.db.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil
}

func (r *LinkPostgres) GetByHash(hash string) (*shortly.Link, error) {
	var link shortly.Link

	result := r.db.Where("hash = ?", hash).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (r *LinkPostgres) GetByID(id uint) (*shortly.Link, error) {
	var link shortly.Link

	result := r.db.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (r *LinkPostgres) GetAll(userID uint, limit, offset int) ([]shortly.Link, error) {
	var links []shortly.Link

	result := r.db.Table("links").
		Where("deleted_at is null AND user_id = ?", userID).
		Session(&gorm.Session{}).
		Order("id").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	if result.Error != nil {
		return nil, result.Error
	}

	return links, nil
}

func (r *LinkPostgres) Count(userID uint) (int64, error) {
	var count int64

	result := r.db.Table("links").
		Where("deleted_at is null AND user_id = ?", userID).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func (r *LinkPostgres) Update(link *shortly.Link, userID uint) (*shortly.Link, error) {
	result := r.db.Model(&shortly.Link{}).Where("id = ? AND user_id = ?", link.ID, userID).Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("link not found or access denied")
	}

	return link, nil
}

func (r *LinkPostgres) Delete(userID, linkID uint) error {
	result := r.db.Where("id = ? AND user_id = ?", linkID, userID).Delete(&shortly.Link{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("link not found or access denied")
	}

	return nil
}
