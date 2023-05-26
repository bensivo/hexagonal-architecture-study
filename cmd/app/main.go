package main

import (
	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders/adapters"
)

func main() {
	repo := adapters.NewInMemoryOrderRepo()
	svc := orders.NewOrderService(repo)
	httpAdapter := adapters.NewHttpAdapter(svc)
	httpAdapter.Start()
}
