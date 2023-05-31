package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func Init(sugar *zap.SugaredLogger, url string) *pgx.Conn {
	sugar.Infof("Connecting to postgres at %s", url)
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		sugar.Error("Failed connecting to postgres")
		sugar.Panic(err) // NOTE: in a real application, we would retry this connection a few times before triggering the panic
	}

	err = RunMigrations(conn, sugar)
	if err != nil {
		sugar.Error("Failed running migrations")
		sugar.Panic(err)
	}

	return conn
}
