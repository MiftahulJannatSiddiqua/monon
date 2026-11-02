package product

import (
	"ecommerce/domain"
)

type Service interface {
	Create(domain.ProductType) (*domain.ProductType, error)
	Get(id int) (*domain.ProductType, error)
	List(page, limit int64) ([]*domain.ProductType, error)
	Count() (int64, error)
	Update(domain.ProductType) (*domain.ProductType, error)
	Delete(id int) error
}
