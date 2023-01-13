package lib

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler with Gin engine initialised
func NewRequestHandler(logger Logger, env Env) RequestHandler {

	// Sentry Initialization
	if env.Environment != "local" && env.SentryDSN != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         env.SentryDSN,
			Environment: `api-gateway-` + env.Environment,
		}); err != nil {
			logger.Infof("Sentry initialization failed: %v\n", err)
		}
	}
	gin.DefaultWriter = logger.GetGinLogger()
	appEnv := env.Environment
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	httpRouter := gin.Default()
	// Attach sentry middleware
	httpRouter.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	return RequestHandler{Gin: httpRouter}
}
