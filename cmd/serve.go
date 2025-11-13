package cmd

import (
	cartDomain "ecommerce/cart"
	"ecommerce/config"
	"ecommerce/infra/db"
	productDomain "ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	cartHandler "ecommerce/rest/handlers/cart"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
	userDomain "ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	config := config.GetConfig()
	dbConnection, err := db.NewDbConnection(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDb(dbConnection, "./migrations")
	if err != nil {
		fmt.Println("Could not run migrations:", err)
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddlewares(config)

	userRepo := repo.NewUserRepo(dbConnection)
	productRepo := repo.NewProductRepo(dbConnection)
	cartRepo := repo.NewCartRepo(dbConnection)

	userService := userDomain.NewUserService(userRepo)
	productService := productDomain.NewProductService(productRepo)
	cartService := cartDomain.NewCartService(cartRepo)

	userHandler := userHandler.NewUserHandler(config, userService)
	productHandler := productHandler.NewProductHandler(middlewares, productService)
	cartHandler := cartHandler.NewCartHandler(middlewares, cartService)

	server := rest.NewServer(
		config,
		userHandler,
		productHandler,
		cartHandler,
	)

	server.Start()
}
