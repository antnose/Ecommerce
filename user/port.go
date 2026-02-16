package user

import (
	"github.com/antnose/Ecommerce/domain"
	userHandler "github.com/antnose/Ecommerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
