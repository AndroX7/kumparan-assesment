package usecase

import (
	"fmt"
	"log"
	"strings"

	"github.com/AndroX7/kumparan-assesment/app/api/middleware"
)

func (u *Usecase) FlushCustomSet(groupSet string, customFieldValue string) {
	groupKey := fmt.Sprint("services:set:", groupSet)

	path := strings.ReplaceAll(groupSet, "_", "-")

	// slug
	for _, routeGroup := range u.routeGroups {

		requestURI := fmt.Sprint(groupSet, ":", routeGroup, "/", path, "/", customFieldValue)
		itemKey := fmt.Sprint(groupSet, ":", requestURI)

		err := u.redis.Delete(middleware.RedisResponsePrefix, itemKey)
		if err != nil {
			log.Print(err)
		}

		err = u.redis.SRem(middleware.RedisResponsePrefix, groupKey, requestURI)
		if err != nil {
			log.Print(err)
		}

		groupKey := fmt.Sprint("services:set:", groupSet, ":", routeGroup, "/", path, "/", customFieldValue)
		requestURIs, _ := u.redis.SMembers(middleware.RedisResponsePrefix, groupKey)

		for _, requestURI := range requestURIs {
			itemKey := fmt.Sprint(groupSet, ":", requestURI)
			_ = u.redis.Delete(middleware.RedisResponsePrefix, itemKey)
			_ = u.redis.SRem(middleware.RedisResponsePrefix, groupKey, requestURI)
		}

	}
}
