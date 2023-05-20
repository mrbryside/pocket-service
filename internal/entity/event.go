package entity

import "github.com/google/uuid"

type Event struct {
	Id        uuid.UUID
	AllEvents []interface{}
}
