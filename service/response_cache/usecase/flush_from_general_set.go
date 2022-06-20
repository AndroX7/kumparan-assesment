package usecase

import (
	"fmt"

	"github.com/AndroX7/kumparan-assesment/app/api/middleware"
)

func (u *Usecase) FlushGeneralSet(groupSet string) {
	groupKey := fmt.Sprint("services:set:", groupSet)
	requestURIs, _ := u.redis.SMembers(middleware.RedisResponsePrefix, groupKey)

	for _, requestURI := range requestURIs {
		itemKey := fmt.Sprint(groupSet, ":", requestURI)

		members, _ := u.redis.SMembers(middleware.RedisResponsePrefix, fmt.Sprint("services:set:", itemKey))

		for _, member := range members {
			itemKey := fmt.Sprint(groupSet, ":", member)
			_ = u.redis.Delete(middleware.RedisResponsePrefix, itemKey)
			_ = u.redis.SRem(middleware.RedisResponsePrefix, groupKey, member)
		}

		_ = u.redis.Delete(middleware.RedisResponsePrefix, itemKey)
		_ = u.redis.SRem(middleware.RedisResponsePrefix, groupKey, requestURI)
	}
}
