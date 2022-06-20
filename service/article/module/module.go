package module

import (
	artistHTTP "github.com/AndroX7/kumparan-assesment/service/article/delivery/http"
	artistRepository "github.com/AndroX7/kumparan-assesment/service/article/repository/mysql"
	artistUsecase "github.com/AndroX7/kumparan-assesment/service/article/usecase"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		artistHTTP.New,
		artistUsecase.New,
		artistRepository.New,
	),
)
