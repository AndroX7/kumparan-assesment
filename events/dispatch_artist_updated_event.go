package events

import (
	"log"

	"github.com/AndroX7/kumparan-assesment/events/entity"
	"github.com/AndroX7/kumparan-assesment/models"

	"github.com/gookit/event"
)

func (e *Event) DispatchArtistUpdatedEvent(artistM *models.Articles) error {
	customEvent := &entity.ArtistEventEntity{ArtistM: artistM}
	customEvent.SetName(EventNameArtistUpdated)
	err := event.FireEvent(customEvent)
	if err != nil {
		log.Printf("error while dispatching %s event: %s\n", EventNameArtistUpdated, err)
		return err
	}

	return nil
}
