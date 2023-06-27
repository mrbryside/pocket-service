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

func (s *Saver) AddTransaction(t entity.Transaction) {
	s.transactions = append(s.transactions, t)
}

func (s *Saver) AddPocketCreatedEvent(pocketId uuid.UUID) error {
	if pocketId == uuid.Nil {
		return errors.New("error add pocket created event")
	}
	evt := entity.EventPocketCreated{
		EventId:   uuid.New(),
		EventType: "pocketCreated",
		PocketId:  pocketId,
		SaverId:   uuid.New(),
		Name:      s.pocket.Name,
		Icon:      s.pocket.Icon,
	}
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
