package eventgen

import "github/mrbryside/pocket-service/internal/entity"

type EventGen struct {
	Pocket       entity.Pocket        // root entity
	Transactions []entity.Transaction // value object
}

func NewEventGen() *EventGen {
	return &EventGen{
		Pocket:       entity.Pocket{},
		Transactions: make([]entity.Transaction, 0),
	}
}

func (e *EventGen) AddTransactions(t entity.Transaction) {
	e.Transactions = append(e.Transactions, t)
}
