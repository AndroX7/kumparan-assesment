package subscriber

import "github.com/AndroX7/kumparan-assesment/listener/event_listeners"

type Client interface {
	SubscribedEvents() map[string]interface{}
}

type Subscriber struct {
	eventListener event_listeners.Client
}

func New(
	eventListener event_listeners.Client,
) Client {
	return &Subscriber{
		eventListener: eventListener,
	}
}
