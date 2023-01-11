package main

import (
	"context"
	"log"

	"api-gateway/api/controllers"
	"api-gateway/api/middlewares"
	"api-gateway/api/routes"
	"api-gateway/client"
	"api-gateway/lib"
	"api-gateway/server"

	"go.uber.org/fx"
)

// bootstrap() : Dependency injection bootstrap function which provide grpcClient, controllers, routes, lib and middelware
// and invoke gin based http server.
func bootstrap() error {
	logger := lib.GetLogger()
	app := fx.New(
		client.GRPCModule,
		controllers.Module,
		routes.Module,
		lib.Module,
		middlewares.Module,
		fx.Options(
			// fx.WithLogger(func() fxevent.Logger {
			// 	return logger.GetFxLogger()
			// }),
			fx.Invoke(
				server.StartHTTPServer,
			)),
		// TODO: figure this out - not initializing with zap logger
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
	app.Run()

	return nil
}

// main : entry point
func main() {
	if err := bootstrap(); err != nil {
		log.Panic(err)
	}
}
