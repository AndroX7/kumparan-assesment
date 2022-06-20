package article

import (
	"github.com/AndroX7/kumparan-assesment/lib/request_util"
	"github.com/AndroX7/kumparan-assesment/models"

	"gorm.io/gorm"
)

type Repository interface {
	Delete(article *models.Articles, tx *gorm.DB) error
	FindByID(articleID uint64) (*models.Articles, error)
	FindAll(config request_util.PaginationConfig) ([]models.Articles, error)
	Insert(article *models.Articles, tx *gorm.DB) error
	Update(article *models.Articles, tx *gorm.DB) error
	Count(config request_util.PaginationConfig) (int64, error)
}
