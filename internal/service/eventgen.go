package service

import (
	"github/mrbryside/pocket-service/internal/domain/eventgen"

	"github.com/google/uuid"
)

type IEventGenService interface {
	GetPocketById(pocketId uuid.UUID) (*eventgen.EventGen, error)
}

type eventGenService struct {
	eventGenRepo eventgen.Repository
}

func NewEventGenService(er eventgen.Repository) IEventGenService {
	return eventGenService{
		eventGenRepo: er,
	}
}

func (eg eventGenService) GetPocketById(pocketId uuid.UUID) (*eventgen.EventGen, error) {
	egAgg, err := eg.eventGenRepo.FindPocketById(pocketId)
	if err != nil {
		return nil, err
	}

	return egAgg, nil
}
