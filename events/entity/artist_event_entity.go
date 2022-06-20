package entity

import (
	"github.com/AndroX7/kumparan-assesment/models"

	"github.com/gookit/event"
)

type ArtistEventEntity struct {
	event.BasicEvent
	ArtistM *models.Articles
}
