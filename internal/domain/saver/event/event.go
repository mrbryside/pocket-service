package event

import "github/mrbryside/pocket-service/internal/domain/saver"

type operation struct{}

func NewOperation() saver.Operation {
	return operation{}
}

func (o operation) InsertPocket(s *saver.Saver) error {
	err := s.AddPocketCreatedEvent(s.PocketEntity().Id)
	if err != nil {
		return err
	}
	return nil
}

func (o operation) InsertTransaction(s *saver.Saver) error {
	s.AddTransactionAddedEvent()
	return nil
}
