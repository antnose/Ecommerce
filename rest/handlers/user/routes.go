package user

import (
	"net/http"

	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Create user
	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
		),
	)

	// Login user
	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)

}
