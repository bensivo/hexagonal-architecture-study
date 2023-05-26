package adapters

import (
	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/stretchr/testify/mock"
)

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
