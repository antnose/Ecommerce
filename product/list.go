package product

import "github.com/antnose/Ecommerce/domain"

func (svc *service) List() ([]*domain.Product, error) {
	return svc.prdctRepo.List()
}
