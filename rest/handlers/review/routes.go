package review

import (
	"net/http"

	middleware "github.com/antnose/Ecommerce/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /reviews", manager.With(
		http.HandlerFunc(h.GetReviews),
	))
}
