package server

import (
	"fmt"
	"time"

	"github.com/Rikypurnomo/warmup/config"
	"github.com/Rikypurnomo/warmup/internal/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type thisRouter struct {
	Router *gin.Engine
}

func NewRouter() *thisRouter {
	return &thisRouter{
		Router: gin.New(),
	}
}

func (route *thisRouter) RoutersConfig() {
	route.Router.Use(middleware.CatchLogApi())
	route.Router.Use(gin.Recovery())
	route.Router.Use(cors.New(config.Cors))
	route.Router.Use(otelgin.Middleware("kiyora"))
	route.Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	route.Router.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{
			"status":  false,
			"code":    405,
			"message": "Method not implemented",
		})
		c.Abort()
	})

	route.Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"status":  false,
			"code":    404,
			"message": "No route found",
		})
		c.Abort()
	})
	RouterApi(route.Router)
}
