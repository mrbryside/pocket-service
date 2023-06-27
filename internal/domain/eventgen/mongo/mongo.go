package mongo

import (
	"context"
	"github/mrbryside/pocket-service/internal/core/db/mongodb"
	"github/mrbryside/pocket-service/internal/domain/eventgen"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
	pocketCollection mongodb.CollectionWrapper
	dbClient         mongodb.ClientWrapper
}

func NewRepository(pocketColl mongodb.CollectionWrapper, c mongodb.ClientWrapper) Repository {
	return Repository{pocketCollection: pocketColl, dbClient: c}
}

func (r Repository) FindPocketById(pocketId uuid.UUID) (*eventgen.EventGen, error) {
	var pk pocketModel
	filter := bson.M{"pocket_id": pocketId}
	err := r.pocketCollection.FindOne(
		context.Background(),
		filter,
	).Decode(&pk)

	if err != nil {
		log.Println(err.Error())
		return &eventgen.EventGen{}, err
	}
	return toAggregate(pk), nil
}

func toAggregate(pk pocketModel) *eventgen.EventGen {
	egAgg := eventgen.NewEventGen()

	genPocketCreatedEvent(pk.Events, egAgg)
	genPocketTransactionAddedEvent(pk.Events, egAgg)

	return egAgg
}

type transactionModel struct {
	Amount    float32   `bson:"amount"`
	Category  string    `bson:"category"`
	CreatedAt time.Time `bson:"created_at"`
}
type pocketModel struct {
	PocketId uuid.UUID    `bson:"pocket_id"`
	Events   []eventModel `bson:"events"`
}
type eventModel struct {
	EventId      uuid.UUID          `bson:"event_id"`
	EventType    string             `bson:"event_type"`
	PocketId     uuid.UUID          `bson:"pocket_id"`
	SaverId      uuid.UUID          `bson:"saver_id"`
	Name         string             `bson:"name"`
	Icon         string             `bson:"icon"`
	Transactions []transactionModel `bson:"transactions"`
}
