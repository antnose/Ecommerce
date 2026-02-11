package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/antnose/Ecommerce/database"
	"github.com/antnose/Ecommerce/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()

	util.SendData(w, createdUser, http.StatusCreated)
}
