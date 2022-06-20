package event_listeners

import (
	"github.com/AndroX7/kumparan-assesment/app/api/middleware"
	"github.com/AndroX7/kumparan-assesment/events/entity"

	"github.com/gookit/event"
)

func (l *Listener) ArtistUpdatedEventListener(e event.Event) error {
	artistM := e.(*entity.ArtistEventEntity).ArtistM

	l.responseCacheUsecase.FlushFromArtist(artistM)
	l.responseCacheUsecase.FlushGeneralSet(middleware.RedisResponseArtistSet)

	return nil
}
