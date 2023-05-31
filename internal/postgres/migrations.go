package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// A migration represents a one-time SQL query that is run on application startup to alter the schema of the database.
// On application start, we check for un-executed migrations and run them one at a time. This ensures our database schema matches our code.
//
// Once a migration has been defined, it should never be deleted to preserve compatiblity with any systems which have applied that migration.
// Instead, just append a new migration to the list.
type Migration struct {
	name      string
	timestamp int
	query     string
}

// RunMigrations ensures all the defined migrations have been run against the given pg connection
// - Creates the migrations table if needed
// - Finds any existing migrations
// - Runs any migrations which are missing, updating the migrations table
//
// This function is idempotent, subsequent calls against the same connection are safe.
func RunMigrations(conn *pgx.Conn, sugar *zap.SugaredLogger) error {
	// Create migrations table
	sugar.Infof("Initializing migrations")
	rows, err := conn.Query(context.Background(), `
		CREATE TABLE IF NOT EXISTS migrations (
			name text,
			timestamp bigint
		);
	`)
	if err != nil {
		sugar.Error("Failed creating migrations table")
		return err
	}
	rows.Close()

	// Collect existing migrations
	existing_migrations := map[string]byte{}
	rows, err = conn.Query(context.Background(), `
		SELECT * from migrations order by timestamp desc
	`)
	if err != nil {
		sugar.Error("Failed querying migrations")
		return err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			sugar.Error("Failed reading migration")
			return err
		}

		existing_migrations[values[0].(string)] = 1
	}
	rows.Close()

	sugar.Infof("Existing migrations: %v", existing_migrations)

	// Run migrations which were not found in the existing_migrations query
	//
	// Put any new migrations here, make sure to ONLY APPEND to this list.
	migrations := []Migration{
		{
			name:      "Orders",
			timestamp: 1685506038,
			query: `
				CREATE TABLE orders (
					id uuid,
					product text,
					quantity int,
					status text
				)
			`,
		},
	}
	for _, migration := range migrations {
		if existing_migrations[migration.name] == 1 {
			sugar.Infof("Skipping migration: %s", migration.name)
			continue
		}
		sugar.Infof("Running migration: %s - %s", migration.name, migration.query)
		rows, err = conn.Query(context.Background(), migration.query)
		if err != nil {
			sugar.Errorf("Failed running migration %s", migration.name)
			return err
		}
		rows.Close()

		rows, err = conn.Query(
			context.Background(),
			`INSERT INTO migrations (name, timestamp) VALUES (@name, @timestamp)`,
			pgx.NamedArgs{
				"name":      migration.name,
				"timestamp": migration.timestamp,
			},
		)
		if err != nil {
			sugar.Error("Failed running migration %s", migration.name)
			return err
		}
		rows.Close()
	}

	return nil
}
