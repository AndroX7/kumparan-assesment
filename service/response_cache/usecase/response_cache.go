package usecase

import (
	"os"

	"github.com/AndroX7/kumparan-assesment/lib/redis"
	"github.com/AndroX7/kumparan-assesment/service/article"
	"github.com/AndroX7/kumparan-assesment/service/response_cache"
)

type Usecase struct {
	redis            redis.Client
	routeGroups      map[string]string
	artistRepository article.Repository
}

func New(
	artistRepository article.Repository,
) response_cache.Usecase {
	redis := redis.NewClient(redis.Credentials{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}, os.Getenv("APP_ENV"))

	routeGroups := map[string]string{
		"admin":  "/admin",
		"user":   "",
		"public": "/public",
		"server": "/server",
	}

	return &Usecase{
		redis:            redis,
		routeGroups:      routeGroups,
		artistRepository: artistRepository,
	}
}
