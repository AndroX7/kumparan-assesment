package mysql

import (
	"log"

	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/errors"

	"gorm.io/gorm"
)

func (r *Repository) Update(article *models.Articles, tx *gorm.DB) error {
	var db = r.db
	if tx != nil {
		db = tx
	}
	err := db.Save(article).Error
	if err != nil {
		log.Println("error-update-artist:", err)
		return errors.ErrUnprocessableEntity
	}
	return nil
}
