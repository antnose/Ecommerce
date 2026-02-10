package handlers

import (
	"net/http"

	"github.com/antnose/Ecommerce/database"
	"github.com/antnose/Ecommerce/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), http.StatusOK)
}
