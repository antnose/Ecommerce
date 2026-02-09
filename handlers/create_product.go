package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/antnose/Ecommerce/product"
	"github.com/antnose/Ecommerce/util"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var newProduct product.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(product.ProductList) + 1
	product.ProductList = append(product.ProductList, newProduct)

	util.SendData(w, newProduct, http.StatusCreated)
}
