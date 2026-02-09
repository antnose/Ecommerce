package handlers

import (
	"net/http"
	"strconv"

	"github.com/antnose/Ecommerce/database"
	"github.com/antnose/Ecommerce/util"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pID {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}

	util.SendData(w, "Data not found", http.StatusNotFound)

}
