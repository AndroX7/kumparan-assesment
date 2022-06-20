package usecase

import (
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/helpers"
	"strconv"
)

func (u *Usecase) Delete(articleID uint64) error {
	var articleM *models.Articles

	err := helpers.ValidateParams(articleID)
	if err != nil {
		return err
	}
	tx := u.transactionManager.NewTransaction()
	tx.Begin()
	{

		articleM, err = u.articleRepo.FindByID(articleID)
		if err != nil {
			return err
		}

		err = u.articleRepo.Delete(articleM, tx)

		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	prefix := strconv.FormatUint(articleID, 10)
	_ = u.redis.Delete(prefix, "-")

	return nil
}
