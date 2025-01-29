package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"route-planner/pkg/database/migration"
)

func main() {
	app := &cli.App{
		Name:     "bun",
		Commands: commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

var commands = []*cli.Command{
	{
		Name:  "migrate",
		Usage: "migrate database",
		Action: func(c *cli.Context) error {
			err := migration.Migrate()
			if err != nil {
				return err
			}
			return nil
		},
	},
	{
		Name:  "resetdb",
		Usage: "reset database",
		Action: func(c *cli.Context) error {
			err := migration.DropTables()
			if err != nil {
				return err
			}
			err = migration.Migrate()
			if err != nil {
				return err
			}
			return nil
		},
	},
}
