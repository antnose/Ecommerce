package cmd

import (
	"github.com/antnose/Ecommerce/config"
	"github.com/antnose/Ecommerce/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
