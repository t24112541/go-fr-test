package routers

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"gofr.dev/pkg/gofr"
)

func (r *routerResource) redisRoutes() {
	// Register your routes here
	// e.g., r.GET("/example", exampleHandler)

	// greeting := greeting.New()

	r.app.GET("/redis", func(ctx *gofr.Context) (any, error) {
		// Get the value using the Redis instance

		val, err := ctx.Redis.Get(ctx.Context, "greeting").Result()
		println("val:", val)
		if err != nil && !errors.Is(err, redis.Nil) {
			// If the key is not found, we are not considering this an error and returning ""
			return nil, err
		}

		return val, nil
	})
}
