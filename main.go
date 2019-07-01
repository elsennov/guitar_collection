package main

import (
	"github.com/urfave/cli"
	"os"
	"personal/guitar_collection/appcontext"
	"personal/guitar_collection/console"
	"personal/guitar_collection/server"
)

func main() {
	appcontext.Init()

	app := cli.NewApp()
	app.Name = "Guitar Collection"
	app.Usage = "Guitars for everyone!"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP api server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:        "migrate",
			Description: "Run database migrations",
			Action: func(c *cli.Context) error {
				console.StartMigrations()
				return nil
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) error {
				console.RollbackLastMigrations()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
