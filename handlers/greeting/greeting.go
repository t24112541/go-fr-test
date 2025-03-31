package greeting

import "gofr.dev/pkg/gofr"

type greetingResource struct{}

func New() *greetingResource {
	return &greetingResource{}
}

func (r *greetingResource) Greeting(ctx *gofr.Context) (any, error) {
	txt_greeting := "Hello geek!"
	return txt_greeting, nil
}
