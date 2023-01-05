package bootstrap

import (
	"context"

	"github.com/saxenashivang/api-gateway/lib"
	"go.uber.org/fx"
)

func GetApp() *fx.App {
	logger := lib.GetLogger()
	opt := Providers
	opts := fx.Options(
		// fx.WithLogger(logger.GetFxLogger),
		GetInvokersOptions(),
	)
	ctx := context.Background()
	app := fx.New(opt, opts)
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
	return app
}
