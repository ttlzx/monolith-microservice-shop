package intraprocess

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttlzx/monolith-microservice-shop/internal/common/price"
	"github.com/ttlzx/monolith-microservice-shop/internal/shop/domain/products"
)

func TestProductFromDomainProduct(t *testing.T) {
	productPrice := price.NewPriceP(42, "USD")
	domainProduct, err := products.NewProduct("123", "name", "desc", productPrice)
	assert.NoError(t, err)

	p := ProductFromDomainProduct(*domainProduct)

	assert.EqualValues(t, Product{
		"123",
		"name",
		"desc",
		productPrice,
	}, p)
}
