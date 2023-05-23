package order

// Output / driven port
type OrderRepository interface {
	Save(order *Order) error

	GetMany() ([]Order, error)

	GetOne(id string) (*Order, error)
}
