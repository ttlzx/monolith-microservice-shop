package intraprocess

import (
	"github.com/pkg/errors"
	"github.com/ttlinzexin/monolith-microservice-shop/internal/common/price"
	"github.com/ttlinzexin/monolith-microservice-shop/internal/shop/domain/products"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       price.Price
}

func ProductFromDomainProduct(domainProduct products.Product) Product {
	return Product{
		string(domainProduct.ID()),
		domainProduct.Name(),
		domainProduct.Description(),
		domainProduct.Price(),
	}
}

type ProductInterface struct {
	repo products.Repository
}

func NewProductInterface(repo products.Repository) ProductInterface {
	return ProductInterface{repo}
}

func (i ProductInterface) ProductByID(id string) (Product, error) {
	domainProduct, err := i.repo.ByID(products.ID(id))
	if err != nil {
		return Product{}, errors.Wrap(err, "cannot get product")
	}

	return ProductFromDomainProduct(*domainProduct), nil
}
