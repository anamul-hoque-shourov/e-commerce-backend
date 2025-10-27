package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	productDomain "ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	userDomain "ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	config := config.GetConfig()
	dbConnection, err := db.NewConnection(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbConnection, "./migrations")
	if err != nil {
		fmt.Println("Could not run migrations:", err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddlewares(config)

	productRepo := repo.NewProductRepo(dbConnection)
	userRepo := repo.NewUserRepo(dbConnection)

	userService := userDomain.NewService(userRepo)
	productService := productDomain.NewService(productRepo)

	productHandler := productHandler.NewHandler(middlewares, productService)
	userHandler := userHandler.NewHandler(config, userService)

	server := rest.NewServer(
		config,
		productHandler,
		userHandler,
	)

	server.Start()
}
