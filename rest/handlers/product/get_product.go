package product

import (
	"net/http"
	"strconv"

	"github.com/antnose/Ecommerce/util"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Get(pID)

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SendData(w, http.StatusOK, product)

}
