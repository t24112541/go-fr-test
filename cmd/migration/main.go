package main

import (
	"context"

	"github.com/t24112541/go-fr-test/cmd/migration/sub"
	"github.com/t24112541/go-fr-test/datasource"
	"github.com/t24112541/go-fr-test/migrations"
	"github.com/uptrace/bun/migrate"
	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.NewCMD()
	ctx := context.Background()

	conn := datasource.New(app, ctx)
	db := conn.RegisterDatasource()

	bunMigrate := migrate.NewMigrator(db, migrations.Migrations)

	sub.RegisterSubCommand(app, bunMigrate)

	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
