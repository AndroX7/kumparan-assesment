package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type Model struct {
	ID   string `uri:"id"`
	Code string `uri:"code"`
	Slug string `uri:"slug"`
}

// CacheRedisCheck sees if there are any cached responses and returns the cached response if one is available.
func (m *Middleware) CacheRedisCheck(groupSet string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the ignoreCache parameter
		ignoreCache := strings.ToLower(c.Query("ignoreCache")) == "true"
		if ignoreCache || os.Getenv("RESPONSE_CACHE_ENABLED") == "false" || os.Getenv("RESPONSE_CACHE_ENABLED") == "" {
			ignoreCache = true
		}

		hasRequestParams := false
		var model Model
		_ = c.ShouldBindUri(&model)
		if len(c.Params) > 0 && (model.ID != "" || model.Code != "") {
			hasRequestParams = true
		}

		cacheIgnoredPaths := []string{"search"}
		isIgnored := funk.Contains(cacheIgnoredPaths, model.Slug)
		isIgnored = funk.Contains(cacheIgnoredPaths, model.Code)

		claims := jwt.ExtractClaims(c)

		var itemKey string
		var role string
		var userID uint64

		if claims["role"] != nil && claims["id"] != nil {
			userID, _ = strconv.ParseUint(fmt.Sprint(claims["id"]), 10, 64)
			role = fmt.Sprint(claims["role"])

			itemKey = fmt.Sprint(groupSet, ":", c.Request.RequestURI, ":user_role:", role, ":user_id:", userID)
		} else {
			itemKey = fmt.Sprint(groupSet, ":", c.Request.RequestURI)
		}

		if !ignoreCache && !isIgnored && role != "admin" {
			groupKey := fmt.Sprint("services:set:", groupSet)

			if c.Request.Method == http.MethodGet {
				// See if we have a cached response
				data := m.redis.Get(RedisResponsePrefix, itemKey)

				if data != "" {
					// If so, use it
					c.Data(200, "application/json", []byte(data))
					c.Abort()
				} else {
					// If not, pass our cache writer to the next middleware
					bcw := &bodyCacheWriter{
						requestURI:       c.Request.RequestURI,
						hasRequestParams: hasRequestParams,
						store:            "redis",
						redis:            m.redis,
						itemKey:          itemKey,
						groupKey:         groupKey,
						authUserRole:     &role,
						authUserID:       &userID,
						model:            &model,
						ResponseWriter:   c.Writer,
					}
					c.Writer = bcw
					c.Next()
				}
			} else {
				groupKey := fmt.Sprint("services:set:", groupSet)
				requestURIs, _ := m.redis.SMembers(RedisResponsePrefix, groupKey)

				for _, requestURI := range requestURIs {
					itemKey := fmt.Sprint(groupSet, ":", requestURI)
					go m.redis.Delete(RedisResponsePrefix, itemKey)
				}

				if len(c.Params) > 0 {
					path := strings.ReplaceAll(groupSet, "_", "-")

					var model Model
					err := c.ShouldBindUri(&model)

					if err != nil {
						return
					} else {
						if model.ID != "" {
							for _, routeGroup := range m.routeGroups {
								requestURI := fmt.Sprint(groupSet, ":", routeGroup, "/", path, "/", model.ID)
								itemKey := fmt.Sprint(groupSet, ":", requestURI)
								go m.redis.Delete(RedisResponsePrefix, itemKey)
								go m.redis.SRem(RedisResponsePrefix, groupKey, requestURI)
							}
						}
						if model.Code != "" {
							for _, routeGroup := range m.routeGroups {
								requestURI := fmt.Sprint(groupSet, ":", routeGroup, "/", path, "/", model.Code)
								itemKey := fmt.Sprint(groupSet, ":", requestURI)
								go m.redis.Delete(RedisResponsePrefix, itemKey)
								go m.redis.SRem(RedisResponsePrefix, groupKey, requestURI)
							}
						}
						if model.Slug != "" {
							for _, routeGroup := range m.routeGroups {
								requestURI := fmt.Sprint(groupSet, ":", routeGroup, "/", path, "/", model.Slug)
								itemKey := fmt.Sprint(groupSet, ":", requestURI)
								go m.redis.Delete(RedisResponsePrefix, itemKey)
								go m.redis.SRem(RedisResponsePrefix, groupKey, requestURI)
							}
						}

						groupKey := fmt.Sprint("services:set:", groupSet, ":", c.Request.RequestURI)
						requestURIs, _ := m.redis.SMembers(RedisResponsePrefix, groupKey)

						for _, requestURI := range requestURIs {
							itemKey := fmt.Sprint(groupSet, ":", requestURI)
							go m.redis.Delete(RedisResponsePrefix, itemKey)
							go m.redis.SRem(RedisResponsePrefix, groupKey, requestURI)
						}
					}
				}

				go m.redis.Delete(RedisResponsePrefix, itemKey)
			}
		}
	}
}
