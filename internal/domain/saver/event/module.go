package event

import (
	"github/mrbryside/pocket-service/internal/domain/saver"

	"go.uber.org/fx"
)

var SaverDomainModule = fx.Module("SaverDomainModule",
	fx.Provide(func() saver.Operation {
		return NewOperation()
	}),
)
