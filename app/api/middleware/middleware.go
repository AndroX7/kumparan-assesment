package middleware

import (
	"log"
	"time"

	"github.com/AndroX7/kumparan-assesment/lib/redis"

	jwt "github.com/appleboy/gin-jwt/v2"
)

const RedisResponseDefaultKeyExpirationTime = 10 * time.Minute
const RedisResponseArtistSet = "article"
const RedisResponsePrefix = "response:article:"

type Config struct {
	// put middleware config here
	JwtToken string
}

type Middleware struct {
	config         Config
	authMiddleware *jwt.GinJWTMiddleware
	redis          redis.Client
	routeGroups    map[string]string
}

func New(
	cfg Config,
	redis redis.Client,
) *Middleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:              []byte(cfg.JwtToken),
		Timeout:          6 * time.Hour,
		MaxRefresh:       6 * time.Hour,
		TimeFunc:         time.Now,
		SigningAlgorithm: "HS512",
	})
	if err != nil {
		log.Fatal("jwt-error:" + err.Error())
	}

	routeGroups := map[string]string{
		"admin":  "/admin",
		"user":   "",
		"public": "/public",
		"server": "/server",
	}

	return &Middleware{
		config:         cfg,
		authMiddleware: authMiddleware,
		redis:          redis,
		routeGroups:    routeGroups,
	}
}
