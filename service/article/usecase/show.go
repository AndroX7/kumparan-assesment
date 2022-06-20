package usecase

import (
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/helpers"
)

func (u *Usecase) Show(articleID uint64) (*models.Articles, error) {

	// try to avoid sql injection by injection query using single quotes
	err := helpers.ValidateParams(articleID)
	if err != nil {
		return nil, err
	}

	article, err := u.articleRepo.FindByID(articleID)
	if err != nil {
		return nil, err
	}

	return article, nil
}
