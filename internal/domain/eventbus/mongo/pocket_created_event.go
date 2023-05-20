package mongo

import (
	"github/mrbryside/pocket-service/internal/domain/eventbus"
	"github/mrbryside/pocket-service/internal/entity"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) insertPocketCreatedEvent(evtAgg *eventbus.EventBus, sc mongo.SessionContext) error {
	icpModel := toInternalCreatedPocketModel(evtAgg)
	if len(icpModel) > 0 {
		_, err := r.pocketCollection.InsertMany(sc, icpModel)
		if err != nil {
			return err
		}
	}
	return nil
}

type CreatedPocketEventModel struct {
	EventId   uuid.UUID `bson:"event_id"`
	EventType string    `bson:"event_type"`
	PocketId  uuid.UUID `bson:"pocket_id"`
	SaverId   uuid.UUID `bson:"saver_id"`
	Name      string    `bson:"name"`
	Icon      string    `bson:"icon"`
}

type CreatedPocketModel struct {
	PocketId uuid.UUID                 `bson:"pocket_id"`
	Events   []CreatedPocketEventModel `bson:"events"`
}

func toInternalCreatedPocketModel(evtAgg *eventbus.EventBus) []interface{} {
	var cpModels []interface{}
	for _, evt := range evtAgg.AllEvents() {
		if val, ok := evt.(entity.EventPocketCreated); ok {
			var evtModels []CreatedPocketEventModel
			evtModel := CreatedPocketEventModel{
				EventId:   val.EventId,
				EventType: val.EventType,
				PocketId:  val.PocketId,
				SaverId:   val.SaverId,
				Name:      val.Name,
				Icon:      val.Icon,
			}
			evtModels = append(evtModels, evtModel)
			cpModel := CreatedPocketModel{
				PocketId: val.PocketId,
				Events:   evtModels,
			}
			cpModels = append(cpModels, cpModel)
		}
	}
	return cpModels
}
