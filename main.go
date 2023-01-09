package main

import (
	"context"
	"fmt"

	"github.com/saxenashivang/api-gateway/api/controllers"
	"github.com/saxenashivang/api-gateway/api/middlewares"
	"github.com/saxenashivang/api-gateway/api/routes"
	"github.com/saxenashivang/api-gateway/client/grpc"
	"github.com/saxenashivang/api-gateway/http"
	"github.com/saxenashivang/api-gateway/lib"
	"github.com/saxenashivang/api-gateway/servers/rest"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// Version is the build version. It is auto assigned at build time.
var Version = ""

func run() error {
	logger := lib.GetLogger()
	app := fx.New(
		grpc.Module,
		controllers.Module,
		routes.Module,
		http.Module,
		middlewares.Module,
		lib.Module,
		fx.Options(
			fx.WithLogger(func() fxevent.Logger {
				return logger.GetFxLogger()
			}),
		),
		fx.Invoke(
			rest.Run,
		),
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
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
