package main

import (
	"context"
	"log"

	"github.com/saxenashivang/api-gateway/api/controllers"
	"github.com/saxenashivang/api-gateway/api/middlewares"
	"github.com/saxenashivang/api-gateway/api/routes"
	"github.com/saxenashivang/api-gateway/client"
	"github.com/saxenashivang/api-gateway/lib"
	"github.com/saxenashivang/api-gateway/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// Dependency injection bootstrap function
func bootstrap() error {
	logger := lib.GetLogger()
	app := fx.New(
		client.GRPCModule,
		controllers.Module,
		routes.Module,
		middlewares.Module,
		lib.Module,
		fx.Options(
			fx.WithLogger(func() fxevent.Logger {
				return logger.GetFxLogger()
			}),
		),
		fx.Invoke(
			server.StartHTTPServer,
		),
		// TODO: figure this out
		// fx.Decorate(logger.GetFxLogger()),
	)

	ctx := context.Background()
	err := app.Start(ctx)
	defer func() {
		err = app.Stop(ctx)
		if err != nil {
			logger.Fatal(err)
		}
	}()
	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Err(); err != nil {
		return err
	}
	app.Run()

	return nil
}

// main : entry point
func main() {
	if err := bootstrap(); err != nil {
		log.Panic(err)
	}
}
