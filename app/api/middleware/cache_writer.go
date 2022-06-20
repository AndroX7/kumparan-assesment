package middleware

import (
	"fmt"

	"github.com/AndroX7/kumparan-assesment/lib/redis"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"

	"github.com/thoas/go-funk"
)

// bodyCacheWriter is used to cache responses in gin.
type bodyCacheWriter struct {
	gin.ResponseWriter
	requestURI       string
	hasRequestParams bool
	store            string
	groupKey         string
	itemKey          string
	authUserRole     *string
	authUserID       *uint64
	redis            redis.Client
	cache            *cache.Cache
	model            *Model
}

// Write a JSON response to gin and cache the response.
func (w bodyCacheWriter) Write(b []byte) (int, error) {

	exceptionPaths := []string{"search"}
	hasException := funk.Contains(exceptionPaths, w.model.Slug)
	hasException = funk.Contains(exceptionPaths, w.model.Code)

	cacheIgnoredPaths := []string{"search"}
	isIgnored := funk.Contains(cacheIgnoredPaths, w.model.Slug)
	isIgnored = funk.Contains(cacheIgnoredPaths, w.model.Code)

	// Write the response to the cache only if a success code
	status := w.Status()
	if 200 <= status && status <= 299 && !isIgnored {
		switch w.store {
		case "redis":
			go w.redis.Set(RedisResponsePrefix, w.itemKey, string(b), RedisResponseDefaultKeyExpirationTime)

			var setKey string
			setKey = w.groupKey
			if w.hasRequestParams && !hasException {
				setKey = fmt.Sprint(w.groupKey, ":with_params")
			}

			itemKeyIndex := w.requestURI
			authRole := *w.authUserRole
			authUserID := *w.authUserID

			if authRole != "" && authUserID != 0 {
				itemKeyIndex = fmt.Sprint(w.requestURI, ":user_role:", authRole, ":user_id:", authUserID)

				if w.model.ID != "" || w.model.Code != "" || (w.model.Slug != "" && !hasException) {
					go w.redis.SAdd(RedisResponsePrefix, fmt.Sprint(w.groupKey, ":", w.requestURI), itemKeyIndex)
				} else {
					if !w.hasRequestParams {
						go w.redis.SAdd(RedisResponsePrefix, setKey, itemKeyIndex)
					}
				}
			}
			go w.redis.SAdd(RedisResponsePrefix, setKey, w.requestURI)
			go w.redis.SAdd(RedisResponsePrefix, fmt.Sprint(w.groupKey, ":all"), w.requestURI)
		}

	}

	// Then write the response to gin
	return w.ResponseWriter.Write(b)
}
