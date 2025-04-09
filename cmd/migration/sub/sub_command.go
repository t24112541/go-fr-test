package sub

import (
	"fmt"
	"strings"

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
	subCMD.rollbackMigration()
	subCMD.makeSQLMigration()
	subCMD.makeGoMigration()
	subCMD.statusMigration()
	subCMD.fakeMigration()
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
		gofr.AddDescription("migration table"),
		gofr.AddHelp("migration table"),
	)
}

func (r *subCommandResource) rollbackMigration() {
	r.app.SubCommand("rollback", func(c *gofr.Context) (any, error) {
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
	},
		gofr.AddDescription("rollback migration"),
		gofr.AddHelp("rollback migration"),
	)
}

func (r *subCommandResource) makeSQLMigration() {
	r.app.SubCommand("make-sql", func(c *gofr.Context) (any, error) {
		name := strings.Join(c.Params("name"), "_")
		files, err := r.migrator.CreateSQLMigrations(c.Context, name)
		if err != nil {
			str := printStatus("make SQL migration file:", err)
			return str, nil
		}

		for _, file := range files {
			fmt.Printf("created migration %s (%s)\n", file.Name, file.Path)
		}

		return nil, nil
	},
		gofr.AddDescription("initial sql migration files"),
		gofr.AddHelp("--name=[migration_name]"))
}

func (r *subCommandResource) makeGoMigration() {
	r.app.SubCommand("make-go", func(c *gofr.Context) (any, error) {
		name := strings.Join(c.Params("name"), "_")
		file, err := r.migrator.CreateGoMigration(c.Context, name)
		if err != nil {
			str := printStatus("make Go migration file:", err)
			return str, nil
		}

		fmt.Printf("created migration %s (%s)\n", file.Name, file.Path)

		return nil, nil
	},
		gofr.AddDescription("initial go migration fies"),
		gofr.AddHelp("--name=[migration_name]"))
}

func (r *subCommandResource) statusMigration() {
	r.app.SubCommand("status", func(c *gofr.Context) (any, error) {
		status, err := r.migrator.MigrationsWithStatus(c.Context)
		if err != nil {
			str := printStatus("status migration:", err)
			return str, nil
		}
		fmt.Printf("migrations: %s\n", status)
		fmt.Printf("unapplied migrations: %s\n", status.Unapplied())
		fmt.Printf("last migration group: %s\n", status.LastGroup())

		return nil, nil
	},
		gofr.AddDescription("initial go migration fies"),
		gofr.AddHelp("--name=[migration_name]"))
}

func (r *subCommandResource) fakeMigration() {
	r.app.SubCommand("fake", func(c *gofr.Context) (any, error) {
		group, err := r.migrator.Migrate(c.Context, migrate.WithNopMigration())
		if err != nil {
			str := printStatus("fake migration:", err)
			return str, nil
		}

		if group.IsZero() {
			err := fmt.Errorf("there are no new migrations to mark as applied\n")
			str := printStatus("fake migration:", err)
			return str, nil
		}

		fmt.Printf("fake as applied %s\n", group)
		str := printStatus("fake migration:", err)

		return str, nil

	},
		gofr.AddDescription("initial go migration fies"),
		gofr.AddHelp("--name=[migration_name]"))
}

func printStatus(statusOf string, errs ...error) string {
	for _, err := range errs {
		if err != nil {
			status := color.RedString("FAILED")

			return fmt.Sprintf("\n%s: %s \ndetails: %s \n", statusOf, status, err)
		}

	}

	status := color.GreenString("OK")
	return fmt.Sprintf("\n%s: %s \n", statusOf, status)
}
