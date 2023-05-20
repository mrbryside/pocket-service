package entity

import "github.com/google/uuid"

type Pocket struct {
	Id   uuid.UUID
	Name string
	Icon string
}
