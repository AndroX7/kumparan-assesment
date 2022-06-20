package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	articleHTTP "github.com/AndroX7/kumparan-assesment/service/article/delivery/http"
	"github.com/AndroX7/kumparan-assesment/service/article/delivery/http/request"
	"github.com/gin-gonic/gin"
)

func TestCreate(t *testing.T) {
	r := gin.Default()
	var h *articleHTTP.Handler

	r.POST("/article", h.Create)

	articleRequest := request.ArticleCreateRequest{
		Title:    "Lorenzo",
		Author:   "My Author",
		Genre:    "Horror",
		ImageUrl: "https://google.com/somewhere-else",
		Created:  "2022-06020",
		Price:    50000,
		Body:     "Lorem Ipsum",
	}
	jsonBody, _ := json.Marshal(articleRequest)
	req, _ := http.NewRequest("POST", "/article", bytes.NewBuffer(jsonBody))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
