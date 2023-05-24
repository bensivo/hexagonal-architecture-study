package orders

// Output port - defines function calls this module makes to external services
// in this case, calls related to order storage and persistence
type OrderRepository interface {
	Save(order *Order) error

	GetMany() ([]Order, error)

	GetOne(id string) (*Order, error)
}
