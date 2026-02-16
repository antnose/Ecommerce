package cmd

import (
	"fmt"
	"os"

	"github.com/antnose/Ecommerce/config"
	"github.com/antnose/Ecommerce/infra/db"
	"github.com/antnose/Ecommerce/repo"
	"github.com/antnose/Ecommerce/rest"
	prdcthandler "github.com/antnose/Ecommerce/rest/handlers/product"
	usrHandler "github.com/antnose/Ecommerce/rest/handlers/user"
	middleware "github.com/antnose/Ecommerce/rest/middlewares"
	"github.com/antnose/Ecommerce/user"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	usrSvc := user.NewService(userRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prdcthandler.NewHandler(middlewares, productRepo)

	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()

}
