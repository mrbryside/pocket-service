package mongo

import (
	"github/mrbryside/pocket-service/internal/core/db/mongodb"
	mockMongoClientWrapper "github/mrbryside/pocket-service/internal/core/generated/mockgen/mongo_client_wrapper"
	mockMongoCollectionWrapper "github/mrbryside/pocket-service/internal/core/generated/mockgen/mongo_collection_wrapper"
	"github/mrbryside/pocket-service/internal/domain/eventbus"
	"github/mrbryside/pocket-service/internal/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertEventsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mockMongoClientWrapper.NewMockClientWrapper(ctrl)
	mockClient.EXPECT().StartSession(gomock.Any()).Return(mongodb.NewMockSession(), nil)

	mockColl := mockMongoCollectionWrapper.NewMockCollectionWrapper(ctrl)
	mockColl.EXPECT().InsertMany(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)

	repo := NewRepository(mockColl, mockClient)
	evtAgg := eventbus.NewEventBus()

	eventCreated := entity.EventPocketCreated{
		EventId:   uuid.New(),
		EventType: "pocketCreated",
		PocketId:  uuid.New(),
		SaverId:   uuid.New(),
		Name:      "test",
		Icon:      "test",
	}

	evtAgg.AddEvent(eventCreated)

	err := repo.InsertEvents(evtAgg)
	assert.Nil(t, err)
}
