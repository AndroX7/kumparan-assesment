package mysql

import (
	"log"

	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/errors"

	"github.com/AndroX7/kumparan-assesment/lib/request_util"
)

func (r *Repository) FindAll(config request_util.PaginationConfig) ([]models.Articles, error) {
	results := []models.Articles{}
	err := r.db.
		Scopes(config.MetaScopes()...).
		Scopes(config.Scopes()...).
		Find(&results).Error
	if err != nil {
		log.Println("error-find-all-artists:", err)
		return nil, errors.ErrUnprocessableEntity
	}
	return results, nil
}
