package orders

type OrderStatus string

const (
	RECEIVED  OrderStatus = "RECEIVED"
	SHIPPED   OrderStatus = "SHIPPED"
	DELIVERED OrderStatus = "DELIVERED"
)

type Order struct {
	ID       string
	Product  string
	Quantity int
	Status   OrderStatus
}
