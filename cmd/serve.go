package cmd

import (
	"fmt"
	"os"

	"github.com/antnose/Ecommerce/config"
	"github.com/antnose/Ecommerce/infra/db"
	"github.com/antnose/Ecommerce/repo"
	"github.com/antnose/Ecommerce/rest"
	"github.com/antnose/Ecommerce/rest/handlers/product"
	"github.com/antnose/Ecommerce/rest/handlers/user"
	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

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
