package mysql

import (
	"log"

	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) Insert(artist *models.Articles, tx *gorm.DB) error {
	var db = r.db

	if tx != nil {
		db = tx
	}
	err := db.Create(artist).Error
	if err != nil {
		log.Println("error-insert-artist:", err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
