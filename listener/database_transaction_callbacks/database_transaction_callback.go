package database_transaction_callbacks

import (
	"github.com/AndroX7/kumparan-assesment/events"
	"github.com/AndroX7/kumparan-assesment/service/article"
	"github.com/AndroX7/kumparan-assesment/service/response_cache"

	"gorm.io/gorm"
)

type Client interface {
	FlushResponseCache(db *gorm.DB)
	FlushRedisKey(db *gorm.DB)
}

type Callback struct {
	db                   *gorm.DB
	events               events.Client
	responseCacheUsecase response_cache.Usecase
	articleUsecase       article.Usecase
}

func New(
	db *gorm.DB,
	events events.Client,
	responseCacheUsecase response_cache.Usecase,
	articleUsecase article.Usecase,
) Client {
	return &Callback{
		db:                   db,
		events:               events,
		responseCacheUsecase: responseCacheUsecase,
		articleUsecase:       articleUsecase,
	}
}
