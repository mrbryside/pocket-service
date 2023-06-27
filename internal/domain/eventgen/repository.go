package eventgen

import "github.com/google/uuid"

type Repository interface {
	FindPocketById(id uuid.UUID) (*EventGen, error)
}
