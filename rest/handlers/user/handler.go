package user

import (
	"github.com/antnose/Ecommerce/config"
	"github.com/antnose/Ecommerce/repo"
)

type Handler struct {
	cnf      *config.Config
	userRepo repo.UserRepo
}

func NewHandler(cnf *config.Config, userRepo repo.UserRepo) *Handler {
	return &Handler{
		cnf:      cnf,
		userRepo: userRepo,
	}

}
