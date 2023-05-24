package orders_test

import (
	"fmt"
	"testing"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// START - mock order repository using testify's mocking library
// TODO - use the vektra/mockery package to auto-generate mock code
type MockOrderRepo struct {
	mock.Mock
}

func (m *MockOrderRepo) Save(order *orders.Order) error {
	args := m.Called(order)
	return args.Error(0)
}
func (m *MockOrderRepo) GetMany() ([]orders.Order, error) {
	args := m.Called()
	return args.Get(0).([]orders.Order), args.Error(1)
}

func (m *MockOrderRepo) GetOne(id string) (*orders.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*orders.Order), args.Error(1)
	}
}

var _ orders.OrderRepository = (*MockOrderRepo)(nil)

// END - mock order repository using testify's mocking library

func TestCreateOrder(t *testing.T) {
	repo := &MockOrderRepo{}
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

	repo := &MockOrderRepo{}
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
	repo := &MockOrderRepo{}
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
	repo := &MockOrderRepo{}
	repo.Mock.On("GetOne", "id").Return(nil, fmt.Errorf("error"))
	s := orders.NewOrderService(repo)

	_, err := s.GetOrder("id")
	assert.EqualError(t, err, "error")
}
