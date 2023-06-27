package mongo

import (
	"github/mrbryside/pocket-service/internal/core/db/mongodb"
	"github/mrbryside/pocket-service/internal/domain/eventbus"

	"go.uber.org/fx"
)

var EventBusRepoModule = fx.Module("EventBusRepoModule",
	fx.Provide(func(pc mongodb.CollectionWrapper, c mongodb.ClientWrapper) eventbus.Repository {
		return NewRepository(pc, c)
	}),
)
