package usecase

import (
	"encoding/json"
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/service/article/delivery/http/request"
	"github.com/AndroX7/kumparan-assesment/utils/helpers"
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
)

func (u *Usecase) Update(request request.ArticleUpdateRequest, articleID uint64) (*models.Articles, error) {

	// try to avoid sql injection by injection query using single quotes
	err := helpers.ValidateParams(articleID)
	if err != nil {
		return nil, err
	}

	articleM, err := u.articleRepo.FindByID(articleID)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(articleM, &request)
	if err != nil {
		log.Println("error-for-copy-request-to-product")
		return nil, err
	}

	err = u.articleRepo.Update(articleM, nil)
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(articleM)
	prefix := strconv.FormatUint(articleM.ID, 10)

	err = u.redis.Set(prefix, "-", string(b), 10*time.Minute)
	if err != nil {
		log.Println("err-set-cache-on-redis: ", err)
	}
	return articleM, err
}
