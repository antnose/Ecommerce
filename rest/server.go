package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/antnose/Ecommerce/config"
	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func Start(cnf config.Config) {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	InitRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}
