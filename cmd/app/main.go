package main

import (
	"fmt"
	"log"

	"github.com/bensivo/hexagonal-architecture-study/internal/domain/orders"
)

func main() {
	orderService := orders.NewOrdersService()

	for i := 0; i < 5; i++ {
		_, err := orderService.CreateOrder("Cheerios", 1)
		if err != nil {
			log.Fatal(err)
		}
	}

	orders, err := orderService.GetOrders()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(orders)

	last, err := orderService.GetOrder(orders[len(orders)-1].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(last)
}
