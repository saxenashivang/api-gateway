package rest

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/saxenashivang/api-gateway/api/middlewares"
	"github.com/saxenashivang/api-gateway/api/routes"
	"github.com/saxenashivang/api-gateway/http/router"
	"github.com/saxenashivang/api-gateway/lib"
)

func Run(
	middleware middlewares.Middlewares,
	env lib.Env,
	route routes.Routes,
	logger lib.Logger,
	router router.Router,
) {
	logger.Info(`+-----------------------+`)
	logger.Info(`| API-GATEWAY |`)
	logger.Info(`+-----------------------+`)

	// Using time zone as specified in env file
	loc, _ := time.LoadLocation(env.TimeZone)
	time.Local = loc

	middleware.Setup()
	route.Setup()

	if env.Environment != "local" && env.SentryDSN != "" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              env.SentryDSN,
			AttachStacktrace: true,
		})
		if err != nil {
			logger.Error("sentry initialization failed")
			logger.Error(err.Error())
		}
	}
	logger.Info("Running server")
	if env.ServerPort == "" {
		if err := router.Run(); err != nil {
			logger.Fatal(err)
			return
		}
	} else {
		if err := router.Run(":" + env.ServerPort); err != nil {
			logger.Fatal(err)
			return
		}
	}
}
