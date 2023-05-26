package orders_test

import (
	"fmt"
	"testing"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/bensivo/hexagonal-architecture-study/internal/orders/adapters"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	repo := &adapters.MockOrderRepo{}
	repo.Mock.On("Save", mock.Anything).Return(nil)

	s := orders.NewOrderService(repo)

	order, err := s.CreateOrder("product", 1)
	if err != nil {
		t.Fail()
		t.Log("Failed to create order")
	}

	_, err = uuid.Parse(order.ID)
	if err != nil {
		t.Fail()
		t.Log("Order id is not a UUID")
	}
}

func TestGetOrders(t *testing.T) {
	res := []orders.Order{}

	repo := &adapters.MockOrderRepo{}
	repo.Mock.On("GetMany").Return(res, nil)

	s := orders.NewOrderService(repo)

	orders, err := s.GetOrders()
	if err != nil {
		t.Fail()
		t.Log("Failed to get orders")
	}

	assert.Equal(t, orders, res, "Equal")
}

func TestGetOrder(t *testing.T) {
	repo := &adapters.MockOrderRepo{}
	repo.Mock.On("GetOne", "id").Return(&orders.Order{}, nil)
	s := orders.NewOrderService(repo)

	ret, err := s.GetOrder("id")
	if err != nil {
		t.Fail()
		t.Log("Failed to create order")
	}

	assert.Equal(
		t,
		&orders.Order{},
		ret,
		"should return order by id",
	)
}

func TestGetOrder404(t *testing.T) {
	repo := &adapters.MockOrderRepo{}
	repo.Mock.On("GetOne", "id").Return(nil, fmt.Errorf("error"))
	s := orders.NewOrderService(repo)

	_, err := s.GetOrder("id")
	assert.EqualError(t, err, "error")
}
