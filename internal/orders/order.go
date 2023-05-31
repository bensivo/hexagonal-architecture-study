package orders

type OrderStatus string

const (
	RECEIVED  OrderStatus = "RECEIVED"
	SHIPPED   OrderStatus = "SHIPPED"
	DELIVERED OrderStatus = "DELIVERED"
)

// Order is a pure domain object, representing an order managed by our application.
//
// NOTE: this struct has no relationships with any specific storage mechism or database. Each implementation of OrderRepository
// will have to convert from this struct to data-types appropriate for the storage mechanism.
type Order struct {
	ID       string
	Product  string
	Quantity int
	Status   OrderStatus
}
