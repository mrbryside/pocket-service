package service

import (
	"errors"
	"testing"

	mockEventBusDomain "github/mrbryside/pocket-service/internal/core/generated/mockgen/eventbus_domain"
	mockSaverDomain "github/mrbryside/pocket-service/internal/core/generated/mockgen/saver_domain"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaverService(t *testing.T) {
	testCases := []struct {
		desc       string
		mockFunc   func(*mockSaverDomain.MockOperation, *mockEventBusDomain.MockRepository)
		pocketId   uuid.UUID
		pocketName string
		pocketIcon string
		error      bool
	}{
		{
			desc: "Test Create Pocket Success should return error=nil",
			mockFunc: func(mockSo *mockSaverDomain.MockOperation, mockEr *mockEventBusDomain.MockRepository) {
				mockSo.EXPECT().InsertPocket(gomock.Any()).Return(nil)
			},
			pocketId:   uuid.New(),
			pocketName: "bank pocket",
			pocketIcon: "this is icon",
			error:      false,
		},
		{
			desc: "Test Create Pocket init aggregate failed should return error",
			mockFunc: func(mockSo *mockSaverDomain.MockOperation, mockEr *mockEventBusDomain.MockRepository) {
			},
			pocketId:   uuid.New(),
			pocketName: "",
			pocketIcon: "this is icon",
			error:      true,
		},
		{
			desc: "Test Create Pocket add created event failed should return error",
			mockFunc: func(mockSo *mockSaverDomain.MockOperation, mockEr *mockEventBusDomain.MockRepository) {
			},
			pocketId: func() uuid.UUID {
				uuid, _ := uuid.Parse("")
				return uuid
			}(),
			pocketName: "bank pocket",
			pocketIcon: "this is icon",
			error:      true,
		},
		{
			desc: "Test Create Pocket insert domain failed should return error",
			mockFunc: func(mockSo *mockSaverDomain.MockOperation, mockEr *mockEventBusDomain.MockRepository) {
				mockSo.EXPECT().InsertPocket(gomock.Any()).Return(errors.New("insert failed"))
			},
			pocketId:   uuid.New(),
			pocketName: "bank pocket",
			pocketIcon: "this is icon",
			error:      true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// Arrange
			ctrl := gomock.NewController(t)
			mockOp := mockSaverDomain.NewMockOperation(ctrl)
			mockRepo := mockEventBusDomain.NewMockRepository(ctrl)
			tC.mockFunc(mockOp, mockRepo)

			// Act
			ss := NewSaverService(mockOp, mockRepo)
			err := ss.InsertPocket(tC.pocketId, tC.pocketName, tC.pocketIcon)

			// Assert
			if tC.error {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
