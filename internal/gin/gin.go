package gin

import (
	"fmt"
	"log"
	"time"

	g "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// createGinLogger returns a gin middleware function
// which logs some basic information about each request using the zap logger
func createGinLogger(sugar *zap.SugaredLogger) g.HandlerFunc {
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

func New(sugar *zap.SugaredLogger) *g.Engine {
	engine := g.New()

	engine.Use(createGinLogger(sugar))
	// engine.Use(g.Logger())
	engine.Use(g.Recovery())
	return engine
}

func Start(engine *g.Engine) {
	fmt.Println("Listening at http://localhost:9999")
	err := engine.Run(":9999")
	if err != nil {
		log.Panicf("Error creating HTTP service: %v\n", err)
	}
}
