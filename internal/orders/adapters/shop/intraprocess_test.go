package shop_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttlzx/monolith-microservice-shop/internal/common/price"
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/domain/orders"
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/infrastructure/shop"
	"github.com/ttlzx/monolith-microservice-shop/internal/shop/interfaces/private/intraprocess"
)

func TestOrderProductFromShopProduct(t *testing.T) {
	shopProduct := intraprocess.Product{
		"123",
		"name",
		"desc",
		price.NewPriceP(42, "EUR"),
	}
	orderProduct, err := shop.OrderProductFromIntraprocess(shopProduct)
	assert.NoError(t, err)

	expectedOrderProduct, err := orders.NewProduct("123", "name", price.NewPriceP(42, "EUR"))
	assert.NoError(t, err)

	assert.EqualValues(t, expectedOrderProduct, orderProduct)
}
