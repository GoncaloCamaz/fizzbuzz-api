/*
Package database provides a wrapper around the Bun ORM for PostgreSQL.
*/
package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// NewDB returns a new bun.DB
func NewDB(c *Config) *bun.DB {
	var connector *pgdriver.Connector

	connector = pgdriver.NewConnector(pgdriver.WithDSN(c.GetDSN()))

	db := sql.OpenDB(connector)

	db.SetMaxOpenConns(c.PoolSize)
	db.SetMaxIdleConns(c.PoolSize)

	bunDB := bun.NewDB(db, pgdialect.New())

	return bunDB
}
