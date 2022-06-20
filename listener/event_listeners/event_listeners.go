package event_listeners

import (
	"github.com/AndroX7/kumparan-assesment/service/response_cache"

	"github.com/gookit/event"
)

type Client interface {
	ArtistUpdatedEventListener(e event.Event) error
}

type Listener struct {
	responseCacheUsecase response_cache.Usecase
}

func New(
	responseCacheUsecase response_cache.Usecase,
) Client {
	return &Listener{
		responseCacheUsecase: responseCacheUsecase,
	}
}
