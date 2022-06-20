package http

import (
	"github.com/AndroX7/kumparan-assesment/app/api/middleware"
	"github.com/AndroX7/kumparan-assesment/service/response_cache"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	responseCacheUsecase response_cache.Usecase
}

func New(referralRewardUC response_cache.Usecase) *Handler {
	return &Handler{
		responseCacheUsecase: referralRewardUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {

}
