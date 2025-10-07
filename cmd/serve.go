package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Serve() {
	config := config.GetConfig()
	rest.Start(config)
}
