package sub

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/uptrace/bun/migrate"
	"gofr.dev/pkg/gofr"
)

type subCommandResource struct {
	app      *gofr.App
	migrator *migrate.Migrator
}

func RegisterSubCommand(app *gofr.App, migrator *migrate.Migrator) {

	subCMD := &subCommandResource{
		app:      app,
		migrator: migrator,
	}

	subCMD.initMigration()
	subCMD.upMigration()

}

func (r *subCommandResource) initMigration() {
	r.app.SubCommand("init", func(c *gofr.Context) (any, error) {
		if err := r.migrator.Init(c.Context); err != nil {
			str := printStatus("initial migration table", err)

			return str, nil
		}

		str := printStatus("initial migration table")
		return str, nil
	},
		gofr.AddDescription("initial migration table"),
		gofr.AddHelp("help initialed migration table"),
	)
}

func (r *subCommandResource) upMigration() {
	r.app.SubCommand("up", func(c *gofr.Context) (any, error) {
		if err := r.migrator.Lock(c.Context); err != nil {
			str := printStatus("migrate", err)

			return str, nil
		}
		defer r.migrator.Unlock(c.Context) //nolint:errcheck

		group, err := r.migrator.Migrate(c.Context)
		if err != nil {
			str := printStatus("migrate", err)

			return str, nil
		}

		if group.IsZero() {
			err := fmt.Errorf("there are no new migrations to run (database is up to date)")
			str := printStatus("migrate", err)

			return str, nil
		}

		str := printStatus("migrate", err)
		return str, nil
	},
		gofr.AddDescription("initial migration table"),
		gofr.AddHelp("help initialed migration table"),
	)
}

func (r *subCommandResource) rollbackMigration() {
	r.app.SubCommand("up", func(c *gofr.Context) (any, error) {
		if err := r.migrator.Lock(c.Context); err != nil {
			str := printStatus("rollback", err)

			return str, nil
		}
		defer r.migrator.Unlock(c.Context) //nolint:errcheck

		group, err := r.migrator.Rollback(c.Context)
		if err != nil {
			str := printStatus("rollback", err)

			return str, nil
		}

		if group.IsZero() {
			err := fmt.Errorf("there are no groups to roll back")
			str := printStatus("rollback", err)

			return str, nil
		}

		str := printStatus("rollback", err)
		return str, nil
	})
}

func printStatus(statusOf string, err ...error) string {
	if err != nil {
		status := color.RedString("FAILED")

		return fmt.Sprintf("\n%s: %s \ndetails: %s \n", statusOf, status, err)
	}

	status := color.GreenString("OK")
	return fmt.Sprintf("\n%s: %s \n", statusOf, status)
}
