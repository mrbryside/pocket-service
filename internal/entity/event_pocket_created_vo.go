package entity

import "github.com/google/uuid"

type EventPocketCreated struct {
	EventId   uuid.UUID
	EventType string
	PocketId  uuid.UUID
	SaverId   uuid.UUID
	Name      string
	Icon      string
}
