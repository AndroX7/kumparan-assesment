package mysql

import (
	"log"

	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) Insert(article *models.Articles, tx *gorm.DB) error {
	var db = r.db

	if tx != nil {
		db = tx
	}
	err := db.Create(article).Error
	if err != nil {
		log.Println("error-insert-article:", err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
