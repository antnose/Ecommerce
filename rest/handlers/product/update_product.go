package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/antnose/Ecommerce/database"
	"github.com/antnose/Ecommerce/util"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product id", http.StatusNotFound)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Please provide a valid JSON", http.StatusBadRequest)
		return
	}

	newProduct.ID = pId

	database.Update(newProduct)

	util.SendData(w, "Successfully updated product", 201)

}
