package mongo

import (
	"github/mrbryside/pocket-service/internal/core/db/mongodb"
	"github/mrbryside/pocket-service/internal/domain/eventbus"
	"github/mrbryside/pocket-service/internal/entity"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPocketEventSuccess(t *testing.T) {
	client := mongodb.NewMongoClient("mongodb+srv://connextor:ConnextorPassword1@connextor.lfncyii.mongodb.net/?retryWrites=true&w=majority")
	clientWrap := mongodb.NewClientWrapper(client)
	db := client.Database("saver-service-dev")
	coll := db.Collection("pockets")
	collWrap := mongodb.Collection{Collection: coll}

	repo := NewRepository(collWrap, clientWrap)

	evtAgg := eventbus.NewEventBus()

	eventCreated := entity.EventPocketCreated{
		EventId:   uuid.New(),
		PocketId:  uuid.New(),
		SaverId:   uuid.New(),
		EventType: "pocketCreated",
		Name:      "test",
		Icon:      "test",
	}
	eventCreatedTwo := entity.EventPocketCreated{
		EventId:   uuid.New(),
		PocketId:  uuid.New(),
		SaverId:   uuid.New(),
		EventType: "pocketCreated",
		Name:      "oh",
		Icon:      "what",
	}
	eventUpdated := entity.EventPocketUpdated{
		EventId:   uuid.New(),
		PocketId:  eventCreated.PocketId,
		EventType: "pocketUpdated",
		Name:      "oh my god",
		Icon:      "what a dog",
	}

	eventAddedTransaction := entity.EventTransactionAdded{
		EventId:      uuid.New(),
		PocketId:     eventCreated.PocketId,
		EventType:    "pocketTransactionAdded",
		Transactions: []entity.Transaction{{200.0, "ซื้อหมู", time.Now()}},
	}
	eventAddedTransactionTwo := entity.EventTransactionAdded{
		EventId:      uuid.New(),
		PocketId:     eventCreated.PocketId,
		EventType:    "pocketTransactionAdded",
		Transactions: []entity.Transaction{{300.0, "ซื้อไก่", time.Now()}},
	}

	evtAgg.AddEvent(eventCreated)
	evtAgg.AddEvent(eventCreatedTwo)
	evtAgg.AddEvent(eventUpdated)
	evtAgg.AddEvent(eventAddedTransaction)
	evtAgg.AddEvent(eventAddedTransactionTwo)

	err := repo.InsertEvents(evtAgg)

	assert.Nil(t, err)
}
