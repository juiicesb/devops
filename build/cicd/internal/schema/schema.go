package schema

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/juiicesb/sqlxmigrator"
)

func Migrate(ctx context.Context, masterDb *sqlx.DB, log *log.Logger, isUnittest bool) error {
	// Load list of Schema migrations and init new sqlxmigrator client
	migrations := migrationList(ctx, masterDb, log, isUnittest)
	m := sqlxmigrator.New(masterDb, sqlxmigrator.DefaultOptions, migrations)
	m.SetLogger(log)

	// Append any schema that need to be applied if this is a fresh migration
	// ie. the migrations database table does not exist.
	m.InitSchema(initSchema(ctx, masterDb, log, isUnittest))

	// Execute the migrations
	return m.Migrate()
}
