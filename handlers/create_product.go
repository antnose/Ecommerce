package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/antnose/Ecommerce/database"
	"github.com/antnose/Ecommerce/util"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, http.StatusCreated)
}
