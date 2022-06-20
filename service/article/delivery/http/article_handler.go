package http

import (
	"github.com/AndroX7/kumparan-assesment/app/api/middleware"
	article "github.com/AndroX7/kumparan-assesment/service/article"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	articleUsecase article.Usecase
}

func New(articleUC article.Usecase) *Handler {
	return &Handler{
		articleUsecase: articleUC,
	}
}

func (h *Handler) Register(r *gin.Engine, m *middleware.Middleware) {
	articleRoute := r.Group("/article")
	{
		articleRoute.GET("", h.Index)
		articleRoute.GET("/:id", h.Show)
		articleRoute.POST("", h.Create)
		articleRoute.DELETE("/:id", h.Delete)
		articleRoute.PUT("/:id", h.Update)
	}
}
