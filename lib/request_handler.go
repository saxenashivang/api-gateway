package lib

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler(logger Logger, env Env) RequestHandler {
	if env.Environment != "local" && env.SentryDSN != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         env.SentryDSN,
			Environment: `clean-backend-` + env.Environment,
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
	return RequestHandler{Gin: httpRouter}
}
