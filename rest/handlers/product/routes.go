package product

import (
	"net/http"

	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Get all products
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(h.GetProducts),
		middleware.Logger,
	))

	// Get single product
	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)

	// Create product
	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProducts),
			h.middlewares.AuthenticateJWT,
		),
	)

	// Update Product
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

	// Delete Product
	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

}
