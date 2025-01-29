package migrations

import (
	"context"
	"route-planner/migration"
	"route-planner/model"

	"github.com/uptrace/bun"
)

func init() {

	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewAddColumn().Model((*Location)(nil)).ColumnExpr("h3_index BIGINT NOT NULL DEFAULT 0").Exec(ctx)
			if err != nil {
				return err
			}
			_, err = tx.NewCreateIndex().
				Model((*Location)(nil)).
				Index("idx_location_h3").
				Column("h3_index").
				Exec(ctx)
			if err != nil {
				return err
			}

			var locations []model.Location
			err = tx.NewSelect().Model(&locations).Scan(ctx)
			if err != nil {
				return err
			}
			for _, loc := range locations {
				loc.SetH3Index()
				_, err = tx.NewUpdate().Model(&loc).Column("h3_index").WherePK().Exec(ctx)
				if err != nil {
					return err
				}
			}

			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewRaw("DROP INDEX idx_location_h3 ON location").Exec(ctx)
			if err != nil {
				return err
			}

			_, err = tx.NewDropColumn().Model((*Location)(nil)).Column("h3_index").Exec(ctx)
			if err != nil {
				return err
			}
			return nil
		})
	}

	migration.Migrations.MustRegister(up, down)
}
