package request

import (
	"github.com/AndroX7/kumparan-assesment/lib/request_util"
)

type ArticleCreateRequest struct {
	Title    string  `json:"title" form:"title" binding:"required"`
	Genre    string  `json:"genre" form:"genre" binding:"required"`
	ImageUrl string  `json:"image_url" form:"image_url" binding:"required"`
	Created  string  `json:"created" form:"created" binding:"required" time_format:"2006-01-02"`
	Price    float64 `json:"price" form:"price" binding:"required"`
	Body     string  `json:"body" form:"body" binding:"required"`
	Author   string  `json:"author" form:"author" binding:"required"`
}

type ArticleUpdateRequest struct {
	Title    *string  `json:"title" form:"title"`
	Genre    *string  `json:"genre" form:"genre"`
	ImageUrl *string  `json:"image_url" form:"image_url"`
	Created  *string  `json:"created" form:"created" time_format:"2006-01-02"`
	Price    *float64 `json:"price" form:"price"`
	Body     string   `json:"body" form:"body" binding:"required"`
	Author   *string  `json:"author" form:"author"`
}

func NewArtistPaginationConfig(conditions map[string][]string) request_util.PaginationConfig {
	request_util.OverrideKey(conditions, "id", "articles.id")
	request_util.OverrideKey(conditions, "created", "articles.created")
	request_util.OverrideKey(conditions, "author", "articles.author")
	request_util.OverrideKey(conditions, "title", "articles.title")
	request_util.OverrideKey(conditions, "body", "articles.body")

	filterable := map[string]string{
		"articles.id":      request_util.IdType,
		"articles.title":   request_util.StringType,
		"articles.author":  request_util.StringType,
		"articles.body":    request_util.StringType,
		"articles.created": request_util.DateType,
	}
	return request_util.NewRequestPaginationConfig(conditions, filterable)
}
