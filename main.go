package main

import (
	"github.com/t24112541/go-fr-test/routers"
	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()

	router := routers.New(app)
	router.RegisterRoutes()

	// Runs the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Run()
}
