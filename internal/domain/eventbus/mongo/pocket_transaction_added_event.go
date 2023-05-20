package mongo

import (
	"github/mrbryside/pocket-service/internal/domain/eventbus"
	"github/mrbryside/pocket-service/internal/entity"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) insertPocketTransactionAddedEvent(e *eventbus.EventBus, sc mongo.SessionContext) error {
	itaEvents := toInternalAddedTransactionEventModel(e)
	if len(itaEvents) > 0 {
		for _, evt := range itaEvents {
			filter := bson.M{"pocket_id": evt.PocketId}
			update := bson.M{"$push": bson.M{"events": evt}}
			_, err := r.pocketCollection.UpdateOne(sc, filter, update)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type TransactionModel struct {
	Amount    float32   `bson:"amount"`
	Category  string    `bson:"category"`
	CreatedAt time.Time `bson:"created_at"`
}

type AddedTransactionEventModel struct {
	EventId      uuid.UUID          `bson:"event_id"`
	EventType    string             `bson:"event_type"`
	PocketId     uuid.UUID          `bson:"pocket_id"`
	Transactions []TransactionModel `bson:"transactions"`
}

func toInternalAddedTransactionEventModel(e *eventbus.EventBus) []AddedTransactionEventModel {
	var evtModels []AddedTransactionEventModel
	for _, evt := range e.AllEvents() {
		if val, ok := evt.(entity.EventTransactionAdded); ok {
			var tms []TransactionModel
			for _, t := range val.Transactions {
				tModel := TransactionModel{
					Amount:    t.Amount,
					Category:  t.Category,
					CreatedAt: t.CreatedAt,
				}
				tms = append(tms, tModel)
			}
			addedEvt := AddedTransactionEventModel{
				EventId:      val.EventId,
				EventType:    val.EventType,
				PocketId:     val.PocketId,
				Transactions: tms,
			}
			evtModels = append(evtModels, addedEvt)
		}
	}
	return evtModels
}
