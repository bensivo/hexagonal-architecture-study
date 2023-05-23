package order

import (
	"github.com/google/uuid"
)

// Input / driving port
type OrdersService interface {
	CreateOrder(product string, quantity int) (*Order, error)

	// ShipOrder(id string) Order

	// DeliverOrder(id string) Order

	GetOrders() ([]Order, error)

	GetOrder(id string) (*Order, error)
}

type OrdersServiceImpl struct {
	repo OrderRepository
}

var _ OrdersService = (*OrdersServiceImpl)(nil)

func NewOrdersService(repo OrderRepository) *OrdersServiceImpl {
	return &OrdersServiceImpl{
		repo: repo,
	}
}

func (os *OrdersServiceImpl) CreateOrder(product string, quantity int) (*Order, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	order := &Order{
		ID:       id.String(),
		Product:  product,
		Quantity: 1,
		Status:   RECEIVED,
	}

	err = os.repo.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrdersServiceImpl) GetOrders() ([]Order, error) {
	orders, err := os.repo.GetMany()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (os *OrdersServiceImpl) GetOrder(id string) (*Order, error) {
	order, err := os.repo.GetOne(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
