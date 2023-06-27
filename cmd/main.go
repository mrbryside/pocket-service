package main

import (
	"context"
	"github/mrbryside/pocket-service/internal/core/db/mongodb"
	"github/mrbryside/pocket-service/internal/domain/eventbus/mongo"
	mongoEv "github/mrbryside/pocket-service/internal/domain/eventgen/mongo"
	"github/mrbryside/pocket-service/internal/domain/saver/event"
	"github/mrbryside/pocket-service/internal/server/echo"
	"github/mrbryside/pocket-service/internal/service"
	"log"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		mongodb.MongoDBModule,
		mongo.EventBusRepoModule,
		event.SaverDomainModule,
		mongoEv.EventGenRepoModule,
		service.SaverServiceModule,
		service.EventGenServiceModule,
		echo.MainHandlerModule,
		echo.EchoServerModule,
		fx.Invoke(run),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	<-app.Done()
}

func run(lifecycle fx.Lifecycle, es echo.EchoServer) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				es.Start()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
