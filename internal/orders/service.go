package orders

import (
	"fmt"

	"github.com/google/uuid"
)

type OrdersServiceImpl struct {
	orders []Order
}

var _ OrdersService = (*OrdersServiceImpl)(nil)

func NewOrdersService() *OrdersServiceImpl {
	return &OrdersServiceImpl{
		orders: make([]Order, 0),
	}
}

func (os *OrdersServiceImpl) CreateOrder(product string, quantity int) (*Order, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	order := Order{
		ID:       id.String(),
		Product:  product,
		Quantity: 1,
		Status:   RECEIVED,
	}

	os.orders = append(os.orders, order)
	return &order, nil
}

func (os *OrdersServiceImpl) GetOrders() ([]Order, error) {
	return os.orders, nil
}

func (os *OrdersServiceImpl) GetOrder(id string) (*Order, error) {
	for _, order := range os.orders {
		if order.ID == id {
			return &order, nil
		}
	}

	return nil, fmt.Errorf("Order %s not found", id)
}
