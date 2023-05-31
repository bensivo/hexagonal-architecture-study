package sqlite

import (
	"database/sql"

	"go.uber.org/zap"
)

type Migration struct {
	name      string
	timestamp int
	query     string
}

func RunMigrations(db *sql.DB, logger *zap.SugaredLogger) error {
	logger.Infof("Initializing migrations table")
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			name text,
			timestamp bigint
		);
	`)
	if err != nil {
		logger.Error("Failed creating migrations table")
		return err
	}

	// Collect existing migrations
	existing_migrations := map[string]byte{}
	rows, err := db.Query("SELECT name from migrations order by timestamp desc")
	if err != nil {
		logger.Error("Failed querying migrations: ", err)
		return err
	}
	for rows.Next() {
		var migration_row struct {
			name string
		}
		err = rows.Scan(&migration_row.name)
		if err != nil {
			logger.Error("Failed reading migration: ", err)
			return err
		}

		existing_migrations[migration_row.name] = 1
	}
	rows.Close()

	// Run migrations which were not found in the existing_migrations query
	//
	// Put any new migrations here, make sure to ONLY APPEND to this list.
	migrations := []Migration{
		{
			name:      "Orders",
			timestamp: 1685506038,
			query: `
				CREATE TABLE orders (
					id text,
					product text,
					quantity int,
					status text
				)
			`,
		},
	}
	for _, migration := range migrations {
		if existing_migrations[migration.name] == 1 {
			logger.Infof("Skipping migration: %s", migration.name)
			continue
		}
		logger.Infof("Running migration: %s - %s", migration.name, migration.query)
		rows, err = db.Query(migration.query)
		if err != nil {
			logger.Errorf("Failed running migration %s", migration.name)
			return err
		}
		rows.Close()

		_, err = db.Exec("INSERT INTO migrations (name, timestamp) VALUES ($1, $2)", migration.name, migration.timestamp)
		if err != nil {
			logger.Error("Failed running migration %s", migration.name)
			return err
		}
		// rows.Close()
	}

	return nil
}
