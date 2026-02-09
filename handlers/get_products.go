package handlers

import (
	"net/http"

	"github.com/antnose/Ecommerce/product"
	"github.com/antnose/Ecommerce/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, product.ProductList, http.StatusOK)

}
