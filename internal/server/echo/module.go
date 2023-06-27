package echo

import (
	"github/mrbryside/pocket-service/internal/service"

	"go.uber.org/fx"
)

var EchoServerModule = fx.Module("EchoServerModule",
	fx.Provide(func(h Handler) EchoServer {
		return NewEchoServer(h)
	}),
)

var MainHandlerModule = fx.Module("MainHandlerModule",
	fx.Provide(func(ss service.ISaverService) Handler {
		return NewHandler(ss)
	}),
)
