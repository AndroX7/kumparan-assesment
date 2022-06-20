package http

import (
	"net/http"

	"github.com/AndroX7/kumparan-assesment/lib/response_util"
	"github.com/AndroX7/kumparan-assesment/service/article/delivery/http/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	articles, articlePagination, err := h.articleUsecase.Index(request.NewArtistPaginationConfig(c.Request.URL.Query()), c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, response_util.IndexResponse{
		Data: articles,
		Meta: articlePagination,
	})
}
