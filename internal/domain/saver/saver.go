package saver

import (
	"errors"
	"github/mrbryside/pocket-service/internal/entity"

	"github.com/google/uuid"
)

type Saver struct {
	pocket       entity.Pocket
	transactions []entity.Transaction
	event        entity.Event
}

func NewSaver(name string, icon string) (*Saver, error) {
	if name == "" || icon == "" {
		return &Saver{}, errors.New("error init saver aggregate")
	}
	pocket := entity.Pocket{Id: uuid.New(), Name: name, Icon: icon}
	transactions := make([]entity.Transaction, 0)
	allEvents := make([]interface{}, 0)

	return &Saver{
		pocket:       pocket,
		transactions: transactions,
		event: entity.Event{
			Id:        uuid.New(),
			AllEvents: allEvents,
		},
	}, nil
}

func NewTransactionSaver(pocketId uuid.UUID) (*Saver, error) {
	if pocketId.String() == "" {
		return &Saver{}, errors.New("error init saver aggregate")
	}

	pocket := entity.Pocket{Id: uuid.New(), Name: "", Icon: ""}
	transactions := make([]entity.Transaction, 0)
	allEvents := make([]interface{}, 0)

	return &Saver{
		pocket:       pocket,
		transactions: transactions,
		event: entity.Event{
			Id:        uuid.New(),
			AllEvents: allEvents,
		},
	}, nil
}
