package cmd

import (
	"fmt"
	"net/http"

	"github.com/antnose/Ecommerce/global_router"
	"github.com/antnose/Ecommerce/handlers"
	"github.com/antnose/Ecommerce/middleware"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	// Get all products
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Logger,
	))

	// Get single product
	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handlers.GetProduct),
		middleware.Logger,
	))

	// Create product
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProducts),
		middleware.Logger,
	))

	fmt.Println("Server running on port: 3000")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
