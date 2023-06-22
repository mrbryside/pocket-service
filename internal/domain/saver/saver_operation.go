package saver

import (
	"errors"
	"github/mrbryside/pocket-service/internal/entity"

	"github.com/google/uuid"
)

func (s *Saver) PocketEntity() entity.Pocket {
	return s.pocket
}

func (s *Saver) EventVo() entity.Event {
	return s.event
}

func (s *Saver) AddPocketCreatedEvent(pocketId uuid.UUID) error {
	if pocketId == uuid.Nil {
		return errors.New("error add pocket created event")
	}
	evt := entity.EventPocketCreated{EventId: uuid.New(), EventType: "pocketCreated", PocketId: pocketId}
	s.event.AllEvents = append(s.event.AllEvents, evt)

	return nil
}

func (s *Saver) AddTransactionAddedEvent() {
	evt := entity.EventTransactionAdded{
		EventId:      uuid.New(),
		EventType:    "transactionAdded",
		PocketId:     s.pocket.Id,
		Transactions: s.transactions,
	}
	s.event.AllEvents = append(s.event.AllEvents, evt)
}
