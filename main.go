package main

import (
	"github.com/bensivo/hexagonal-architecture-study/internal/gin"
	"github.com/bensivo/hexagonal-architecture-study/internal/logging"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders/adapters"
	"github.com/bensivo/hexagonal-architecture-study/internal/postgres"
)

func main() {
	logger := logging.Init()
	engine := gin.New(logger)

	conn := postgres.Init(logger, "postgres://user:password@postgres:5432/order_service")
	repo := adapters.NewPostgresOrderRepo(conn, logger)

	// repo := adapters.NewInMemoryOrderRepo()

	svc := orders.NewOrderService(repo)

	ginAdapter := adapters.NewGinAdapter(svc, logger)
	ginAdapter.RegisterRoutes(engine)

	gin.Start(engine)
}
