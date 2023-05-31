package main

import (
	"github.com/bensivo/hexagonal-architecture-study/internal/gin"
	"github.com/bensivo/hexagonal-architecture-study/internal/logging"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders/adapters"
	"github.com/bensivo/hexagonal-architecture-study/internal/sqlite"
	_ "modernc.org/sqlite"
)

func main() {
	logger := logging.Init()
	engine := gin.New(logger)

	// Here, we're demonstrating the plug-and-play functionality we get with hexagonal architecture.
	// We can inject either the SQLite repo or the Postgres repo into the Order Service, and no code in the order service has to change.
	//  - of course, if using the postgres repo, make sure the postgres docker container is running

	db := sqlite.Init("./data/order_service.db", logger)
	repo := adapters.NewSqliteOrderRepo(db, logger)

	// conn := postgres.Init(logger, "postgres://user:password@postgres:5432/order_service")
	// repo := adapters.NewPostgresOrderRepo(conn, logger)

	svc := orders.NewOrderService(repo)

	ginAdapter := adapters.NewGinAdapter(svc, logger)
	ginAdapter.RegisterRoutes(engine)

	gin.Start(engine)
}
