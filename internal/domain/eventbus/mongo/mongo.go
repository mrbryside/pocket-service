package mongo

import (
	"context"
	"github/mrbryside/pocket-service/internal/core/db/mongodb"
	"github/mrbryside/pocket-service/internal/domain/eventbus"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	pocketCollection mongodb.CollectionWrapper
	dbClient         mongodb.ClientWrapper
}

func NewRepository(pocketColl mongodb.CollectionWrapper, c mongodb.ClientWrapper) Repository {
	return Repository{pocketCollection: pocketColl, dbClient: c}
}

func (r Repository) InsertEvents(e *eventbus.EventBus) error {
	// new session
	session, err := r.dbClient.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	err = session.StartTransaction()
	if err != nil {
		return err
	}
	sc := mongo.NewSessionContext(context.Background(), session)

	// events
	err = r.insertPocketCreatedEvent(e, sc)
	if err != nil {
		return err
	}
	err = r.insertPocketUpdatedEvent(e, sc)
	if err != nil {
		return err
	}
	err = r.insertPocketTransactionAddedEvent(e, sc)
	if err != nil {
		return err
	}

	// commit the transaction if there were no errors
	err = session.CommitTransaction(context.Background())
	if err != nil {
		return err
	}

	return nil
}
