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

	product := database.Get(pID)
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SendData(w, product, http.StatusOK)

}
