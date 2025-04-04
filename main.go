package main

import (
	"context"

	"github.com/t24112541/go-fr-test/datasource"
	"github.com/t24112541/go-fr-test/routers"
	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()
	ctx := context.Background()

	conn := datasource.New(app, ctx)
	_ = conn.RegisterDatasource()

	app.SubCommand("hello", func(c *gofr.Context) (any, error) {
		return "Hello World!", nil
	},
		gofr.AddDescription("Print 'Hello World!'"),
		gofr.AddHelp("hello world option"),
	)

	router := routers.New(app)
	router.RegisterRoutes()

	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
