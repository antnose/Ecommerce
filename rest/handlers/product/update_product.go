package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/antnose/Ecommerce/repo"
	"github.com/antnose/Ecommerce/util"
)

type ReqUpdateProduct struct {
	Title       string  `json:title`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	var req ReqUpdateProduct

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)

	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		ImgUrl:      req.ImgUrl,
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusOK, "Successfully updated product")

}
