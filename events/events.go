package events

import (
	"github.com/AndroX7/kumparan-assesment/models"
)

const (
	EventNameArtistUpdated = "artist_updated"
)

type Client interface {
	DispatchArtistUpdatedEvent(artistM *models.Articles) error
}

type Event struct {
}

func Register() Client {
	return &Event{}
}
