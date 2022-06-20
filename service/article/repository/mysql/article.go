package mysql

import (
	"github.com/AndroX7/kumparan-assesment/models"
	article "github.com/AndroX7/kumparan-assesment/service/article"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type Model struct {
	productM models.Articles
}

func New(
	db *gorm.DB,
) article.Repository {
	return &Repository{
		db: db,
	}
}
