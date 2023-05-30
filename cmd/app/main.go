package main

import (
	"github.com/bensivo/hexagonal-architecture-study/internal/gin"
	"github.com/bensivo/hexagonal-architecture-study/internal/logging"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders/adapters"
)

func main() {
	logger := logging.Init()
	engine := gin.New(logger)

	repo := adapters.NewInMemoryOrderRepo()
	svc := orders.NewOrderService(repo)

	httpAdapter := adapters.NewHttpAdapter(svc, logger)
	httpAdapter.RegisterRoutes(engine)

	gin.Start(engine)
}
