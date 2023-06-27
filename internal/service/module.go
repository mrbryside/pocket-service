package service

import (
	"github/mrbryside/pocket-service/internal/domain/eventbus"
	"github/mrbryside/pocket-service/internal/domain/eventgen"
	"github/mrbryside/pocket-service/internal/domain/saver"

	"go.uber.org/fx"
)

var SaverServiceModule = fx.Module("SaverServiceModule",
	fx.Provide(func(e eventbus.Repository, s saver.Operation) ISaverService {
		return NewSaverService(s, e)
	}),
)

var EventGenServiceModule = fx.Module("EventGenServiceModule",
	fx.Provide(func(er eventgen.Repository) IEventGenService {
		return NewEventGenService(er)
	}),
)
