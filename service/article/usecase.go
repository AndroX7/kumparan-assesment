package article

import (
	"github.com/AndroX7/kumparan-assesment/lib/request_util"
	"github.com/AndroX7/kumparan-assesment/models"
	"github.com/AndroX7/kumparan-assesment/service/article/delivery/http/request"

	"github.com/AndroX7/kumparan-assesment/lib/response_util"

	"github.com/gin-gonic/gin"
)

type Usecase interface {
	Index(paginationConfig request_util.PaginationConfig, c *gin.Context) ([]models.Articles, response_util.PaginationMeta, error)
	Show(articleID uint64) (*models.Articles, error)
	Create(request request.ArticleCreateRequest) (*models.Articles, error)
	Update(request request.ArticleUpdateRequest, articleID uint64) (*models.Articles, error)
	Delete(articleID uint64) error
}
