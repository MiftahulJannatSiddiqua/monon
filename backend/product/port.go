package product

import (
	"ecommerce/domain"
	prdctHandler "ecommerce/rest/handlers/product"
)

type Service interface {
	prdctHandler.Service
}

type ProductRepo interface {
	Create(p domain.ProductType) (*domain.ProductType, error)
	List(page, limit int64) ([]*domain.ProductType, error)
	Count() (int64, error)
	Get(productId int) (*domain.ProductType, error)
	Update(product domain.ProductType) (*domain.ProductType, error)
	Delete(deleteProduct int) error
}
