package schema

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

// Migrate function
func Migrate(db *sql.DB) error {
	driver := darwin.NewGenericDriver(db, darwin.PostgresDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
