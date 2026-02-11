package product

import (
	"encoding/json"
	"net/http"

	"github.com/antnose/Ecommerce/database"
	"github.com/antnose/Ecommerce/util"
)

func (h *Handler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	createdProduct := database.Store(newProduct)
	util.SendData(w, createdProduct, http.StatusCreated)
}
