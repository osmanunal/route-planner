package database

import (
	"fmt"
	"github.com/uptrace/bun/schema"
	"route-planner/pkg/config"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *bun.DB

func ConnectDB(config config.DBConfig) *bun.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.User, config.Password, config.Host, config.Port, config.Name,
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	schema.SetTableNameInflector(func(s string) string {
		return s
	})

	DB = bun.NewDB(sqlDB, mysqldialect.New())
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	DB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return DB
}
