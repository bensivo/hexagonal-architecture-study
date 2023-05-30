package gin

import (
	"time"

	g "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// createLoggingMiddleware returns a gin middleware function
// which logs some basic information about each request using the zap logger
func createLoggingMiddleware(sugar *zap.SugaredLogger) g.HandlerFunc {
	return func(c *g.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		duration := time.Since(start).Microseconds()

		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		sugar.Infof("%s %s %d %dµs", method, path, statusCode, duration)
	}
}
