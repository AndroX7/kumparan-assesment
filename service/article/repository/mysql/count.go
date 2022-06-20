package mysql

import (
	"log"

	"github.com/AndroX7/kumparan-assesment/lib/request_util"
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/errors"
)

func (r *Repository) Count(config request_util.PaginationConfig) (int64, error) {
	var count int64

	err := r.db.
		Model(&models.Articles{}).
		Scopes(config.Scopes()...).
		Count(&count).Error

	if err != nil {
		log.Println("error-count-artist:", err)
		return 0, errors.ErrUnprocessableEntity
	}
	return count, nil
}
