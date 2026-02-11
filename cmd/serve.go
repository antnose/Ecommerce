package cmd

import (
	"github.com/antnose/Ecommerce/config"
	"github.com/antnose/Ecommerce/rest"
	"github.com/antnose/Ecommerce/rest/handlers/product"
	"github.com/antnose/Ecommerce/rest/handlers/review"
	"github.com/antnose/Ecommerce/rest/handlers/user"
	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()

}
