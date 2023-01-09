package router

import (
	"net/http"

	"github.com/saxenashivang/api-gateway/lib"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

// NewRouter : all the routes are defined here
func NewRouter(
	env lib.Env,
	logger lib.Logger,
) Router {

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

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Attach sentry middleware
	httpRouter.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	httpRouter.GET("/health-check", func(c *gin.Context) {
		// utils.SendSentryMsg(c, "Error")
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running"})
	})

	httpRouter.GET("/ping", func(c *gin.Context) {
		// utils.SendSentryMsg(c, "Error")
		c.JSON(http.StatusOK, gin.H{"data": "Pong"})
	})

	return Router{
		httpRouter,
	}
}