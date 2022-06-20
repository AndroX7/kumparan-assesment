package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) BasicHandle() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("BASIC_USERNAME"): os.Getenv("BASIC_PASSWORD"),
	})
}
