package mongo

import "github/mrbryside/pocket-service/internal/domain/eventgen"

func genPocketCreatedEvent(em []eventModel, eg *eventgen.EventGen) {
	for _, val := range em {
		if val.EventType == "pocketCreated" {
			eg.Pocket.Id = val.PocketId
			eg.Pocket.Name = val.Name
			eg.Pocket.Icon = val.Icon
		}
	}
}
