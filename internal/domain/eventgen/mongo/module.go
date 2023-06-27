package mongo

import (
	"github/mrbryside/pocket-service/internal/core/db/mongodb"
	"github/mrbryside/pocket-service/internal/domain/eventgen"

	"go.uber.org/fx"
)

var EventGenRepoModule = fx.Module("EventGenRepoModule",
	fx.Provide(func(pc mongodb.CollectionWrapper, c mongodb.ClientWrapper) eventgen.Repository {
		return NewRepository(pc, c)
	}),
)
