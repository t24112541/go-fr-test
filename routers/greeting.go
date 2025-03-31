package routers

import (
	"gofr.dev/pkg/gofr"
)

func (r *routerResource) greetingRoutes() {
	// Register your routes here
	// e.g., r.GET("/example", exampleHandler)

	// greeting := greeting.New()

	r.app.GET("/", func(ctx *gofr.Context) (any, error) {
		txt_greeting := "Hello geekkkk!"
		return txt_greeting, nil
	})
}
