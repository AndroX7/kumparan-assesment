package usecase

import (
	"fmt"

	"github.com/AndroX7/kumparan-assesment/models"

	"github.com/AndroX7/kumparan-assesment/lib/request_util"
	"github.com/AndroX7/kumparan-assesment/lib/response_util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (u *Usecase) Index(paginationConfig request_util.PaginationConfig, c *gin.Context) ([]models.Articles, response_util.PaginationMeta, error) {
	meta := response_util.PaginationMeta{
		Offset: paginationConfig.Offset(),
		Limit:  paginationConfig.Limit(),
		Total:  0,
	}

	if search, ok := c.Request.URL.Query()["search"]; ok {
		paginationConfig.AddScope(func(db *gorm.DB) *gorm.DB {
			return db.Where("AND title like ? OR author like ? ", fmt.Sprint("%", search[0], "%"), fmt.Sprint("%", search[0], "%"))
		})
	}

	articles, err := u.articleRepo.FindAll(paginationConfig)
	if err != nil {
		return nil, meta, err
	}

	total, err := u.articleRepo.Count(paginationConfig)
	if err != nil {
		return nil, meta, err
	}
	meta.Total = total

	return articles, meta, nil
}
