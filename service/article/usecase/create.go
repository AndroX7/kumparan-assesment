package usecase

import (
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/service/article/delivery/http/request"
)

func (u *Usecase) Create(request request.ArticleCreateRequest) (*models.Articles, error) {
	articleM := &models.Articles{
		Author:   request.Author,
		Title:    request.Title,
		Genre:    request.Genre,
		ImageUrl: request.ImageUrl,
		Created:  request.Created,
		Price:    request.Price,
		Body:     request.Body,
	}
	err := u.articleRepo.Insert(articleM, nil)
	if err != nil {
		return nil, err
	}
	return articleM, nil
}
