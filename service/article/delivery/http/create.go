package http

import (
	"log"
	"net/http"

	"github.com/AndroX7/kumparan-assesment/service/article/delivery/http/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var request request.ArticleCreateRequest

	if err := c.ShouldBind(&request); err != nil {
		log.Println("error-on-create-new-article: ", err)
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	articleM, err := h.articleUsecase.Create(request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, articleM)
}
