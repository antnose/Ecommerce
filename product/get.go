package product

import "github.com/antnose/Ecommerce/domain"

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.prdctRepo.Get(id)
}
