package product

import "github.com/antnose/Ecommerce/domain"

func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Update(prdct)
}
