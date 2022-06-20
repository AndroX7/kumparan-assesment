package usecase

import (
	"encoding/json"
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/service/article/delivery/http/request"
	"github.com/AndroX7/kumparan-assesment/utils/errors"
	"log"
	"strconv"
	"time"
)

func (u *Usecase) Create(request request.ArticleCreateRequest) (*models.Articles, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, request.Created)
	if err != nil {
		customError := errors.ErrUnprocessableEntity
		customError.Message = err.Error()
		return nil, customError
	}
	articleM := &models.Articles{
		Author:   request.Author,
		Title:    request.Title,
		Genre:    request.Genre,
		ImageUrl: request.ImageUrl,
		Created:  t,
		Price:    request.Price,
		Body:     request.Body,
	}
	err = u.articleRepo.Insert(articleM, nil)
	if err != nil {
		log.Println("error-on-create-new-article: ", err)
		return nil, err
	}

	b, _ := json.Marshal(articleM)
	prefix := strconv.FormatUint(articleM.ID, 10)

	err = u.redis.Set(prefix, "-", string(b), 10*time.Minute)
	if err != nil {
		log.Println("err-set-cache-on-redis: ", err)
	}
	return articleM, nil
}
