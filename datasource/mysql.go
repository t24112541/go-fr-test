package datasource

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/t24112541/go-fr-test/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"gofr.dev/pkg/gofr"
)

type datasourceResource struct {
	app *gofr.App
	ctx context.Context
}

func New(app *gofr.App, ctx context.Context) *datasourceResource {
	return &datasourceResource{
		app: app,
		ctx: ctx,
	}
}

func (r *datasourceResource) RegisterDatasource() (db *bun.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		r.app.Config.Get("DB_USER"),
		r.app.Config.Get("DB_PASSWORD"),
		r.app.Config.Get("DB_HOST"),
		r.app.Config.Get("DB_PORT"),
		r.app.Config.Get("DB_NAME"),
	)
	sqldb, err := sql.Open(
		r.app.Config.Get("DB_DIALECT"),
		dsn,
	)
	if err != nil {
		panic(err)
	}

	db = bun.NewDB(sqldb, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	r.RegisterM2MModels(db)

	return
}

func (r *datasourceResource) RegisterM2MModels(db *bun.DB) {
	// Register many-to-many model so Bun can better recognize the m2m relation.
	models := []interface{}{
		&models.UserBook{}, // Ensure this is now valid
	}

	for _, model := range models {
		db.RegisterModel(model)
	}

	// Generate migration for the UserBook model
	// migration, err := GenerateMigrationFromModel(db, &models.UserBook{})
	// if err != nil {
	//     panic(fmt.Sprintf("Failed to generate migration: %v", err))
	// }
	// fmt.Println(migration)
}

// func GenerateMigrationFromModel(db *bun.DB, model interface{}) (string, error) {
//     createTableQuery := db.NewCreateTable().
//         Model(model).
//         String()

//     dropTableQuery := db.NewDropTable().
//         Model(model).
//         IfExists().
//         String()

//     return fmt.Sprintf("-- up\n%s\n\n-- down\n%s", createTableQuery, dropTableQuery), nil
// }
