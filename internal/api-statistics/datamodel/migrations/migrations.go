/*
Package migrations hold postgres migrations
*/
package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

var sqlMigrations embed.FS

// Migrations hold and manage the service's migrations
var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
