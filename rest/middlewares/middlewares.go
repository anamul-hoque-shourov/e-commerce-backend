package middlewares

import "ecommerce/config"

type Middlewares struct {
	config *config.Config
}

func NewMiddlewares(config *config.Config) *Middlewares {
	return &Middlewares{
		config: config,
	}
}
