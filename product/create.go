package product

import "github.com/antnose/Ecommerce/domain"

func (svc *service) Create(prdct domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Create(prdct)
}
