package eventbus

import (
	"github/mrbryside/pocket-service/internal/entity"

	"github.com/google/uuid"
)

type EventBus struct {
	event entity.Event // root entity
}

func NewEventBus() (*EventBus, error) {
	evt := entity.Event{
		Id:        uuid.New(),
		AllEvents: make([]interface{}, 0),
	}

	return &EventBus{
		event: evt,
	}, nil
}

func (e *EventBus) EventEntity() entity.Event {
	return e.event
}

func (e *EventBus) AllEvents() []interface{} {
	return e.event.AllEvents
}

func (e *EventBus) AddEvent(evt interface{}) {
	e.event.AllEvents = append(e.event.AllEvents, evt)
}
