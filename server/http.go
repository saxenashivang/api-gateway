package server

import (
	"time"

	"api-gateway/api/middlewares"
	"api-gateway/api/routes"
	"api-gateway/lib"

	"github.com/getsentry/sentry-go"
)

func StartHTTPServer(
	middleware middlewares.Middlewares,
	env lib.Env,
	router lib.RequestHandler,
	route routes.Routes,
	logger lib.Logger,
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
		if err := router.Gin.Run(); err != nil {
			logger.Fatal(err)
			return
		}
	} else {
		if err := router.Gin.Run(":" + env.ServerPort); err != nil {
			logger.Fatal(err)
			return
		}
	}
}
