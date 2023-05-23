package orders

type OrderStatus string

const (
	RECEIVED  OrderStatus = "RECEIVED"
	SHIPPED   OrderStatus = "SHIPPED"
	DELIVERED OrderStatus = "DELIVERED"
)

type Order struct {
	ID string

	Product  string
	Quantity int
	Status   OrderStatus
}

// Input / driving port
type OrdersService interface {
	CreateOrder(product string, quantity int) (*Order, error)

	// ShipOrder(id string) Order

	// DeliverOrder(id string) Order

	GetOrders() ([]Order, error)

	GetOrder(id string) (*Order, error)
}

// Output / driven port
type OrdersRepository interface {
	Save(order Order) error

	GetMany() ([]Order, error)

	GetOne(id string) (Order, error)
}
