package rest

import (
	"net/http"

	"github.com/antnose/Ecommerce/rest/handlers"
	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Get all products
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Logger,
	))

	// Get single product
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)

	// Create product
	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProducts),
		),
	)

	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
		),
	)

	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
		),
	)

}
