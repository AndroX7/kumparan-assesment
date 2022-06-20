package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	//
	articleHTTP "github.com/AndroX7/kumparan-assesment/service/article/delivery/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	r := gin.Default()
	var h *articleHTTP.Handler
	r.GET("/article?title=Lorem&author=rowling", h.Index)

	req, _ := http.NewRequest("GET", "/article?title=Lorem&author=rowling", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
