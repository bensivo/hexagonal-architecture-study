package storage

import (
	"fmt"

	"github.com/bensivo/hexagonal-architecture-study/internal/domain/order"
)

type InMemoryOrderRepo struct {
	orders map[string]order.Order
}

var _ order.OrderRepository = (*InMemoryOrderRepo)(nil)

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[string]order.Order),
	}
}

func (r *InMemoryOrderRepo) Save(order *order.Order) error {
	r.orders[order.ID] = *order
	return nil
}

func (r *InMemoryOrderRepo) GetMany() ([]order.Order, error) {
	res := make([]order.Order, 0)
	for _, v := range r.orders {
		res = append(res, v)
	}

	return res, nil
}

func (r *InMemoryOrderRepo) GetOne(id string) (*order.Order, error) {
	order, exists := r.orders[id]
	if !exists {
		return nil, fmt.Errorf("Order %s not found", id)
	}
	return &order, nil
}
