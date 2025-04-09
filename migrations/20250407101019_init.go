package migrations

import (
	"context"
	"fmt"

	"github.com/t24112541/go-fr-test/models"
	"github.com/uptrace/bun"
)

// should use for init table only
func init() {
	normalModels := []interface{}{
		(*models.User)(nil),
		(*models.Book)(nil),
	}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) (err error) {
		fmt.Print(" [up migration] ")

		// regis m2m model before create table
		db.RegisterModel((*models.UserBook)(nil))

		for _, model := range normalModels {
			if _, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx); err != nil {
				return err
			}
		}

		if _, err := db.NewCreateTable().
			Model((*models.UserBook)(nil)).
			ForeignKey(`(user_id) REFERENCES users (user_id) ON DELETE CASCADE`).
			ForeignKey(`(book_id) REFERENCES books (book_id) ON DELETE CASCADE`).
			IfNotExists().
			Exec(ctx); err != nil {

			return err
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) (err error) {
		fmt.Print(" [down migration] ")

		db.NewDropTable().Model((*models.UserBook)(nil)).IfExists().Exec(ctx)

		for _, model := range normalModels {
			if _, err := db.NewDropTable().Model(model).IfExists().Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	})
}
