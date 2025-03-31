package routers

import (
	"github.com/t24112541/go-fr-test/models"
	"gofr.dev/pkg/gofr"
)

type routerResource struct {
	app *gofr.App
}

func New(app *gofr.App) *routerResource {
	return &routerResource{
		app: app,
	}
}

func (r *routerResource) RegisterRoutes() {
	// AddRESTHandlers creates CRUD handles for the given entity
	err := r.app.AddRESTHandlers(&models.CustomerEntity{})
	if err != nil {
		return
	}

	// Register your routes here
	r.greetingRoutes()
	r.redisRoutes()
}
