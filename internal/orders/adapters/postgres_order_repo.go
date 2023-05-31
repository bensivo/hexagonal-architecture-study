package adapters

import (
	"context"
	"fmt"

	"github.com/bensivo/hexagonal-architecture-study/internal/orders"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type PostgresOrderRepo struct {
	conn   *pgx.Conn // TODO: we should be using a connection pool here, instead of a single connection
	logger *zap.SugaredLogger
}

var _ orders.OrderRepository = (*PostgresOrderRepo)(nil)

func NewPostgresOrderRepo(conn *pgx.Conn, logger *zap.SugaredLogger) *PostgresOrderRepo {
	return &PostgresOrderRepo{
		conn:   conn,
		logger: logger,
	}
}

func (por *PostgresOrderRepo) Save(order *orders.Order) error {
	rows, err := por.conn.Query(
		context.Background(),
		"INSERT INTO orders (id, product, quantity, status) VALUES (@id, @product, @quantity, @status)",
		pgx.NamedArgs{
			"id":       order.ID,
			"product":  order.Product,
			"quantity": order.Quantity,
			"status":   order.Status,
		},
	)
	if err != nil {
		por.logger.Error("Failed inserting order: ", err)
		return err
	}
	rows.Close()
	return nil
}

func (por *PostgresOrderRepo) GetMany() ([]orders.Order, error) {
	res := []orders.Order{}
	rows, err := por.conn.Query(context.Background(), "SELECT id, product, quantity, status from orders")
	if err != nil {
		por.logger.Error("Failed querying orders")
		return nil, err
	}

	for rows.Next() {
		var order_row struct {
			id       [16]byte // postgres stores UUIDs as byte arrays, not strings
			product  string
			quantity int
			status   string
		}
		err = rows.Scan(&order_row.id, &order_row.product, &order_row.quantity, &order_row.status)
		if err != nil {
			por.logger.Error("Failed scanning row: ", err)
			return nil, err
		}

		res = append(res, orders.Order{
			ID:       fmt.Sprintf("%x-%x-%x-%x-%x", order_row.id[0:4], order_row.id[4:6], order_row.id[6:8], order_row.id[8:10], order_row.id[10:16]),
			Product:  order_row.product,
			Quantity: order_row.quantity,
			Status:   orders.OrderStatus(order_row.status),
		})
	}
	rows.Close()
	return res, nil
}

func (por *PostgresOrderRepo) GetOne(id string) (*orders.Order, error) {
	row := por.conn.QueryRow(context.Background(), "SELECT id, product, quantity, status from orders where id = @id", pgx.NamedArgs{
		"id": id,
	})

	var order_row struct {
		id       [16]byte // postgres stores UUIDs as byte arrays, not strings
		product  string
		quantity int
		status   string
	}
	err := row.Scan(&order_row.id, &order_row.product, &order_row.quantity, &order_row.status)
	if err != nil {
		por.logger.Error("Failed scanning row: ", err)
		return nil, err
	}

	return &orders.Order{
		ID:       fmt.Sprintf("%x-%x-%x-%x-%x", order_row.id[0:4], order_row.id[4:6], order_row.id[6:8], order_row.id[8:10], order_row.id[10:16]),
		Product:  order_row.product,
		Quantity: order_row.quantity,
		Status:   orders.OrderStatus(order_row.status),
	}, nil
}
