package adapters

import (
	"database/sql"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"go.uber.org/zap"
)

type SqliteOrderRepo struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

var _ orders.OrderRepository = (*SqliteOrderRepo)(nil)

func NewSqliteOrderRepo(db *sql.DB, logger *zap.SugaredLogger) *SqliteOrderRepo {
	return &SqliteOrderRepo{
		db:     db,
		logger: logger,
	}
}

func (repo *SqliteOrderRepo) Save(order *orders.Order) error {
	_, err := repo.db.Exec(
		"INSERT INTO orders (id, product, quantity, status) VALUES ($1, $2, $3, $4)",
		order.ID,
		order.Product,
		order.Quantity,
		order.Status,
	)
	if err != nil {
		repo.logger.Error("Failed to insert record into order repository", err)
		return err
	}

	return nil
}

func (repo *SqliteOrderRepo) GetMany() ([]orders.Order, error) {
	res := []orders.Order{}

	rows, err := repo.db.Query("SELECT id, product, quantity, status from orders")
	if err != nil {
		repo.logger.Error("Failed querying orders from repo", err)
		return nil, err
	}

	for rows.Next() {
		var order_row struct {
			id       string
			product  string
			quantity int
			status   string
		}

		err = rows.Scan(&order_row.id, &order_row.product, &order_row.quantity, &order_row.status)
		if err != nil {
			repo.logger.Error("Failed parsing order rows", err)
			return nil, err
		}

		res = append(res, orders.Order{
			ID:       order_row.id,
			Product:  order_row.product,
			Quantity: order_row.quantity,
			Status:   orders.OrderStatus(order_row.status),
		})
	}
	rows.Close()

	return res, nil
}

func (repo *SqliteOrderRepo) GetOne(id string) (*orders.Order, error) {
	row := repo.db.QueryRow("SELECT id, product, quantity, status from orders where id = $1", id)
	var order_row struct {
		id       string
		product  string
		quantity int
		status   string
	}
	err := row.Scan(&order_row.id, &order_row.product, &order_row.quantity, &order_row.status)
	if err != nil {
		repo.logger.Error("Failed parsing order rows", err)
		return nil, err
	}

	return &orders.Order{
		ID:       order_row.id,
		Product:  order_row.product,
		Quantity: order_row.quantity,
		Status:   orders.OrderStatus(order_row.status),
	}, nil
}
