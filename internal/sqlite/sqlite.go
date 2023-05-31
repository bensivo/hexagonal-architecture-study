package sqlite

import (
	"database/sql"

	"go.uber.org/zap"
	_ "modernc.org/sqlite"
)

func Init(filepath string, logger *zap.SugaredLogger) *sql.DB {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		panic(err)
	}

	err = RunMigrations(db, logger)
	if err != nil {
		panic(err)
	}

	return db
}
