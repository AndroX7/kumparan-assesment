package http

import (
	"net/http"
	"strconv"

	"github.com/AndroX7/kumparan-assesment/utils/errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Show(c *gin.Context) {
	articleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		_ = c.Error(errors.ErrUnprocessableEntity).SetType(gin.ErrorTypePublic)
	}

	articleM, err := h.articleUsecase.Show(articleID)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, articleM)
}
