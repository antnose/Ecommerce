package cmd

import (
	"github.com/antnose/Ecommerce/config"
	"github.com/antnose/Ecommerce/rest"
	"github.com/antnose/Ecommerce/rest/handlers/product"
	"github.com/antnose/Ecommerce/rest/handlers/review"
	"github.com/antnose/Ecommerce/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler()
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
