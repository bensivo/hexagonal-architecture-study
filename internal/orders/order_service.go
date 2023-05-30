package orders

import (
	"github.com/google/uuid"
)

// Input / driving port - defines the functions exposed by this application to external components
// Driving adapters should implement this interface to expose this application using a specific protocol
//
//	Example - an HTTP adapter would create HTTP endpoints for each function
//	Example - a GRPC adapter would expose a GRPC service with functions for each service
type OrderService interface {
	CreateOrder(product string, quantity int) (*Order, error)
	GetOrders() ([]Order, error)
	GetOrder(id string) (*Order, error)
}

type OrderServiceImpl struct {
	repo OrderRepository
}

var _ OrderService = (*OrderServiceImpl)(nil)

func NewOrderService(repo OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		repo: repo,
	}
}

func (os *OrderServiceImpl) CreateOrder(product string, quantity int) (*Order, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	order := &Order{
		ID:       id.String(),
		Product:  product,
		Quantity: quantity,
		Status:   RECEIVED,
	}

	err = os.repo.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrderServiceImpl) GetOrders() ([]Order, error) {
	orders, err := os.repo.GetMany()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (os *OrderServiceImpl) GetOrder(id string) (*Order, error) {
	order, err := os.repo.GetOne(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}
