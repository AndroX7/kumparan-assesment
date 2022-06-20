package usecase

import (
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/helpers"
)

func (u *Usecase) Delete(articleID uint64) error {
	var articleM *models.Articles
	// try to avoid sql injection by injection query using single quotes
	err := helpers.ValidateParams(articleID)
	if err != nil {
		return err
	}
	tx := u.transactionManager.NewTransaction()
	tx.Begin()
	{

		articleM, err = u.articleRepo.FindByID(articleID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	err = u.articleRepo.Delete(articleM, nil)

	if err != nil {
		return err
	}

	return nil
}
