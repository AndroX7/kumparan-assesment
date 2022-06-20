package usecase

import (
	"encoding/json"
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/utils/helpers"
	"strconv"
)

func (u *Usecase) Show(articleID uint64) (*models.Articles, error) {
	var article *models.Articles
	err := helpers.ValidateParams(articleID)
	if err != nil {
		return nil, err
	}

	prefix := strconv.FormatUint(articleID, 10)
	data := u.redis.Get(prefix, "-")

	if data != "" {
		_ = json.Unmarshal([]byte(data), &article)
		return article, nil
	}
	article, err = u.articleRepo.FindByID(articleID)
	if err != nil {
		return nil, err
	}

	return article, nil
}
