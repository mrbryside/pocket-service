package service

import (
	"github/mrbryside/pocket-service/internal/domain/eventbus"
	"github/mrbryside/pocket-service/internal/domain/saver"
	"github/mrbryside/pocket-service/internal/entity"
	"time"

	"github.com/google/uuid"
)

type ISaverService interface {
	InsertPocket(pocketName, pocketIcon string) error
	InsertTransaction(pocketId uuid.UUID, amount float32, category string) error
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

func (ss saverService) InsertPocket(pocketName string, pocketIcon string) error {
	saverAgg, err := saver.NewSaver(pocketName, pocketIcon)
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

func (ss saverService) InsertTransaction(pocketId uuid.UUID, amount float32, category string) error {
	saverAgg, err := saver.NewTransactionSaver(pocketId)
	if err != nil {
		return err
	}
	tVo := entity.Transaction{
		Amount:    amount,
		Category:  category,
		CreatedAt: time.Now(),
	}
	saverAgg.AddTransaction(tVo)

	err = ss.saverOperation.InsertTransaction(saverAgg)
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
