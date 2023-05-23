package orders_test

import (
	"testing"

	"github.com/bensivo/hexagonal-architecture-study/internal/domain/orders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	s := orders.NewOrdersService()
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
	s := orders.NewOrdersService()

	for i := 0; i < 5; i++ {
		_, err := s.CreateOrder("product", 1)
		if err != nil {
			t.Fail()
			t.Log("Failed to create order")
		}

	}

	orders, err := s.GetOrders()
	if err != nil {
		t.Fail()
		t.Log("Failed to get orders")
	}

	if len(orders) != 5 {
		t.Fail()
		t.Logf("Expected 5 orders, but got %d", len(orders))
	}
}

func TestGetOrder(t *testing.T) {
	s := orders.NewOrdersService()

	order, err := s.CreateOrder("product", 1)
	if err != nil {
		t.Fail()
		t.Log("Failed to create order")
	}

	ret, err := s.GetOrder(order.ID)
	if err != nil {
		t.Fail()
		t.Log("Failed to create order")
	}

	assert.Equal(
		t,
		order,
		ret,
		"should return order by id",
	)
}

func TestGetOrder404(t *testing.T) {
	s := orders.NewOrdersService()

	_, err := s.CreateOrder("product", 1)
	if err != nil {
		t.Fail()
		t.Log("Failed to create order")
	}

	_, err = s.GetOrder("id")

	assert.EqualError(t, err, "Order id not found")
}
