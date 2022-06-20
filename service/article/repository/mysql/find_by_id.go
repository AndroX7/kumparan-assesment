package mysql

import (
	"log"

	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) FindByID(artistID uint64) (*models.Articles, error) {
	model := models.Articles{}
	err := r.db.
		Where("id = ?", artistID).
		First(&model).Error

	if err == gorm.ErrRecordNotFound || &model == nil {
		return nil, errors.ErrNotFound
	}

	if err != nil {
		log.Println("error-find-artist-by-id:", err)
		return nil, errors.ErrUnprocessableEntity
	}

	return &model, nil
}
