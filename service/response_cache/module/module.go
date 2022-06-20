package module

import (
	"go.uber.org/fx"

	responseCacheUsecase "github.com/AndroX7/kumparan-assesment/service/response_cache/usecase"
)

var Module = fx.Options(
	fx.Provide(
		responseCacheUsecase.New,
	),
)
