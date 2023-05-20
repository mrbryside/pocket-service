package entity

import "github.com/google/uuid"

type EventPocketUpdated struct {
	EventId   uuid.UUID
	EventType string
	PocketId  uuid.UUID
	Name      string
	Icon      string
}
