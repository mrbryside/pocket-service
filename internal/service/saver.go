package service

import (
	"github/mrbryside/pocket-service/internal/domain/eventbus"
	"github/mrbryside/pocket-service/internal/domain/saver"

	"github.com/google/uuid"
)

type ISaverService interface {
	InsertPocket(pocketId uuid.UUID, pocketName, pocketIcon string) error
}

type saverService struct {
	saverOperation saver.Operation
	eventbusRepo   eventbus.Repository
}

func NewSaverService(so saver.Operation, er eventbus.Repository) ISaverService {
	return saverService{
		saverOperation: so,
		eventbusRepo:   er,
	}
}

func (ss saverService) InsertPocket(pocketId uuid.UUID, pocketName string, pocketIcon string) error {
	saverAgg, err := saver.NewSaver(pocketName, pocketIcon)
	if err != nil {
		return err
	}
	err = saverAgg.AddPocketCreatedEvent(pocketId)
	if err != nil {
		return err
	}

	err = ss.saverOperation.InsertPocket(saverAgg)
	if err != nil {
		return err
	}

	eb := eventbus.NewEventBus()
	eb.AddEvents(saverAgg.EventVo().AllEvents)

	err = ss.eventbusRepo.InsertEvents(eb)
	if err != nil {
		return err
	}

	return nil
}
