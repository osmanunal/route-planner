package migrations

import (
	"context"
	"route-planner/migration"
	"route-planner/pkg/model"

	"github.com/uptrace/bun"
)

type Location struct {
	model.BaseModel

	Name      string
	Latitude  float64
	Longitude float64
	Color     string
}

func init() {
	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewCreateTable().Model(&Location{}).IfNotExists().WithForeignKeys().Exec(ctx)
			return err
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewDropTable().Model(&Location{}).IfExists().Exec(ctx)
			return err
		})
	}

	migration.Migrations.MustRegister(up, down)
}
