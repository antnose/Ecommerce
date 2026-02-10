package handlers

import (
	"net/http"
	"strconv"

	"github.com/antnose/Ecommerce/database"
	"github.com/antnose/Ecommerce/util"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please provide a valid id", http.StatusBadRequest)
		return
	}

	database.Delete(pID)

	util.SendData(w, "Successfully delete product", 201)
}
