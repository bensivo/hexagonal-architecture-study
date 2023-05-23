package order_test

import (
	"fmt"
	"testing"

	"github.com/bensivo/hexagonal-architecture-study/internal/domain/order"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// START - mock order repository using testify's mocking library
// TODO - use the vektra/mockery package to auto-generate mock code
type MockOrderRepo struct {
	mock.Mock
}

func (m *MockOrderRepo) Save(order *order.Order) error {
	args := m.Called(order)
	return args.Error(0)
}
func (m *MockOrderRepo) GetMany() ([]order.Order, error) {
	args := m.Called()
	return args.Get(0).([]order.Order), args.Error(1)
}

func (m *MockOrderRepo) GetOne(id string) (*order.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).(*order.Order), args.Error(1)
	}
}

var _ order.OrderRepository = (*MockOrderRepo)(nil)

// END - mock order repository using testify's mocking library

func TestCreateOrder(t *testing.T) {
	repo := &MockOrderRepo{}
	repo.Mock.On("Save", mock.Anything).Return(nil)

	s := order.NewOrdersService(repo)

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
	repo := &MockOrderRepo{}
	repo.Mock.On("GetMany").Return([]order.Order{}, nil)

	s := order.NewOrdersService(repo)

	orders, err := s.GetOrders()
	if err != nil {
		t.Fail()
		t.Log("Failed to get orders")
	}

	assert.Equal(t, orders, []order.Order{}, "Equal")
}

func TestGetOrder(t *testing.T) {
	repo := &MockOrderRepo{}
	repo.Mock.On("GetOne", "id").Return(&order.Order{}, nil)
	s := order.NewOrdersService(repo)

	ret, err := s.GetOrder("id")
	if err != nil {
		t.Fail()
		t.Log("Failed to create order")
	}

	assert.Equal(
		t,
		&order.Order{},
		ret,
		"should return order by id",
	)
}

func TestGetOrder404(t *testing.T) {
	repo := &MockOrderRepo{}
	repo.Mock.On("GetOne", "id").Return(nil, fmt.Errorf("error"))
	s := order.NewOrdersService(repo)

	_, err := s.GetOrder("id")
	assert.EqualError(t, err, "error")
}
