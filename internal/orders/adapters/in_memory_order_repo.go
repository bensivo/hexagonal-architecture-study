package adapters

import (
	"fmt"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
)

type InMemoryOrderRepo struct {
	orders map[string]orders.Order
}

var _ orders.OrderRepository = (*InMemoryOrderRepo)(nil)

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[string]orders.Order),
	}
}

func (r *InMemoryOrderRepo) Save(order *orders.Order) error {
	r.orders[order.ID] = *order
	return nil
}

func (r *InMemoryOrderRepo) GetMany() ([]orders.Order, error) {
	res := make([]orders.Order, 0)
	for _, v := range r.orders {
		res = append(res, v)
	}

	return res, nil
}

func (r *InMemoryOrderRepo) GetOne(id string) (*orders.Order, error) {
	order, exists := r.orders[id]
	if !exists {
		return nil, fmt.Errorf("order %s not found", id)
	}
	return &order, nil
}
