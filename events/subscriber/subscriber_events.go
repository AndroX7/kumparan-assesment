package subscriber

import (
	"github.com/AndroX7/kumparan-assesment/events"

	"github.com/gookit/event"
)

func (s *Subscriber) SubscribedEvents() map[string]interface{} {
	return map[string]interface{}{
		events.EventNameArtistUpdated: event.ListenerFunc(s.eventListener.ArtistUpdatedEventListener),
	}
}
