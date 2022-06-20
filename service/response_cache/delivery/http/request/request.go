package request

type ResponseCacheFlushRequest struct {
	SetName string `form:"set_name" json:"set_name" binding:"required"`
}
