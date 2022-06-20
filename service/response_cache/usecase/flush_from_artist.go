package usecase

import (
	"github.com/AndroX7/kumparan-assesment/app/api/middleware"
	"github.com/AndroX7/kumparan-assesment/models"
)

func (u *Usecase) FlushFromArtist(artistM *models.Articles) {
	groupSet := middleware.RedisResponseArtistSet

	u.FlushGeneralSet(groupSet)
	u.FlushCustomSet(groupSet, artistM.Title)
	u.FlushGeneralSetWithID(groupSet, artistM.ID)
}
