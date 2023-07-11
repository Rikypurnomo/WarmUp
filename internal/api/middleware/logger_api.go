package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// make middleware to get request and response
// Path: api\middleware\logger_api.go
func CatchLogApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		var trace_id = c.GetHeader("X-Request-Id")
		if c.GetHeader("X-Request-Id") == "" {
			trace_id = uuid.New().String()
			c.Header("X-Request-Id", trace_id)
		}
		c.Set("X-Request-Id", trace_id)
		c.Next()
	}
}

func CheckServicesAdapter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if az.IsShutdown && config.IsEnabledAdapterAz() {
		// 	c.JSON(503, gin.H{
		// 		"status":  false,
		// 		"code":    503,
		// 		"message": "Services Adapter is down",
		// 	})
		// 	c.Abort()
		// 	return
		// }
		c.Next()
	}
}
