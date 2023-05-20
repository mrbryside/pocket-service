package entity

import "github.com/google/uuid"

type EventTransactionAdded struct {
	EventId      uuid.UUID
	EventType    string
	PocketId     uuid.UUID
	Transactions []Transaction
}
