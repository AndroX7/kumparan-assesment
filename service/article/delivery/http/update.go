package http

import (
	"net/http"
	"strconv"

	"github.com/AndroX7/kumparan-assesment/service/article/delivery/http/request"
	"github.com/AndroX7/kumparan-assesment/utils/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	var request request.ArticleUpdateRequest

	articleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(errors.ErrUnprocessableEntity).SetType(gin.ErrorTypePublic)
	}

	// validate request
	if err := c.ShouldBind(&request); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	articleM, err := h.articleUsecase.Update(request, articleID)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, articleM)
}
