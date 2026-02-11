package rest

import (
	"net/http"

	"github.com/antnose/Ecommerce/rest/handlers"
	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// #######
	// --------> Products APIS <----------

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

	// Update Product
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
		),
	)

	// Delete Product
	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
		),
	)

	// #######
	// --------> Users APIS <----------

	// Create user
	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		),
	)

	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)

}
