package migration

import (
	"context"
	"database/sql"
	"route-planner/migration"
	"route-planner/pkg/config"
	"route-planner/pkg/database"

	"log"

	"github.com/uptrace/bun"
)

func Migrate() error {
	log.Println("Migration Start")
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	db := database.ConnectDB(cfg.DBConfig)
	models := migration.GetMigrationModels()
	err = db.RunInTx(context.Background(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for _, m := range models {
			if _, err = tx.NewCreateTable().Model(m).IfNotExists().WithForeignKeys().Exec(ctx); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	log.Println("Migration End")
	return nil
}

func DropTables() error {
	log.Println("Drop Tables Start")
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	db := database.ConnectDB(cfg.DBConfig)
	models := migration.GetMigrationModels()
	err = db.RunInTx(context.Background(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		for _, m := range models {
			if _, err = tx.NewDropTable().Model(m).IfExists().Cascade().Exec(ctx); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	log.Println("Drop Tables End")
	return nil
}
