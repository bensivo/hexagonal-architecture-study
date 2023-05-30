package gin

import (
	"log"

	g "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func New(sugar *zap.SugaredLogger) *g.Engine {
	engine := g.New()
	engine.Use(createLoggingMiddleware(sugar))
	engine.Use(g.Recovery()) // Turns panics into 500 status codes
	return engine
}

func Start(engine *g.Engine) {
	err := engine.Run(":9999")
	if err != nil {
		log.Panicf("Error creating HTTP service: %v\n", err)
	}
}
