package main

import (
	"fmt"
	"log"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders/adapters"
)

func main() {
	repo := adapters.NewInMemoryOrderRepo()

	svc := orders.NewOrderService(repo)

	for i := 0; i < 5; i++ {
		_, err := svc.CreateOrder("Cheerios", 1)
		if err != nil {
			log.Fatal(err)
		}
	}

	orders, err := svc.GetOrders()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(orders)

	last, err := svc.GetOrder(orders[len(orders)-1].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(last)
}
