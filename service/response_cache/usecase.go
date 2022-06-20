package response_cache

import "github.com/AndroX7/kumparan-assesment/models"

type Usecase interface {
	FlushFromArtist(artist *models.Articles)

	FlushAllFromSet(groupSet string)
	FlushGeneralSet(groupSet string)
	FlushGeneralSetWithID(groupSet string, ID uint64)
	FlushCustomSet(groupSet string, customFieldValue string)
}
