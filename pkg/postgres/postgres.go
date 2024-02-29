package postgres

import (
	"fmt"
	"ledger-service-pismo/pkg/config"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// ConnectDB Returns new Postgresql db instance
// Param - config object
// Returns - connection instance of postgres sqlx db
// Returns - error if any
func ConnectDB(c *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlPassword,
	)

	db, err := sqlx.Connect(c.Postgres.PgDriver, dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Connect")
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "db.Ping")
	}

	return db, nil
}
