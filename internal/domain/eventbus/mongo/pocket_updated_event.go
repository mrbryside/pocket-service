package mongo

import (
	"github.com/google/uuid"
	"github/mrbryside/pocket-service/internal/domain/eventbus"
	"github/mrbryside/pocket-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) insertPocketUpdatedEvent(evtAgg *eventbus.EventBus, sc mongo.SessionContext) error {
	iupEvents := toInternalUpdatedPocketEventModel(evtAgg)
	if len(iupEvents) > 0 {
		for _, evt := range iupEvents {
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

type UpdatedPocketEventModel struct {
	EventId   uuid.UUID `bson:"event_id"`
	EventType string    `bson:"event_type"`
	PocketId  uuid.UUID `bson:"pocket_id"`
	Name      string    `bson:"name"`
	Icon      string    `bson:"icon"`
}

func toInternalUpdatedPocketEventModel(evtAgg *eventbus.EventBus) []UpdatedPocketEventModel {
	var evtModels []UpdatedPocketEventModel
	for _, evt := range evtAgg.AllEvents() {
		if val, ok := evt.(entity.EventPocketUpdated); ok {
			evtModel := UpdatedPocketEventModel{
				EventId:   val.EventId,
				EventType: val.EventType,
				PocketId:  val.PocketId,
				Name:      val.Name,
				Icon:      val.Icon,
			}
			evtModels = append(evtModels, evtModel)
		}
	}
	return evtModels
}
