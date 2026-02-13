package cmd

import (
	"github.com/antnose/Ecommerce/config"
	"github.com/antnose/Ecommerce/repo"
	"github.com/antnose/Ecommerce/rest"
	"github.com/antnose/Ecommerce/rest/handlers/product"
	"github.com/antnose/Ecommerce/rest/handlers/user"
	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)

	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()

}
