package orders

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttlzx/monolith-microservice-shop/internal/common/price"
)

func TestNewOrder(t *testing.T) {
	orderProduct, orderAddress := createOrderContent(t)

	testOrder, err := NewOrder("1", orderProduct, orderAddress)
	assert.NoError(t, err)

	assert.EqualValues(t, orderProduct, testOrder.Product())
	assert.EqualValues(t, orderAddress, testOrder.Address())
	assert.False(t, testOrder.Paid())
}

func TestNewOrder_empty_id(t *testing.T) {
	orderProduct, orderAddress := createOrderContent(t)

	_, err := NewOrder("", orderProduct, orderAddress)
	assert.EqualValues(t, ErrEmptyOrderID, err)
}

func TestOrder_MarkAsPaid(t *testing.T) {
	orderProduct, orderAddress := createOrderContent(t)

	testOrder, err := NewOrder("1", orderProduct, orderAddress)
	assert.NoError(t, err)

	assert.False(t, testOrder.Paid())
	testOrder.MarkAsPaid()
	assert.True(t, testOrder.Paid())
}

func createOrderContent(t *testing.T) (Product, Address) {
	productPrice, err := price.NewPrice(10, "USD")
	assert.NoError(t, err)

	orderProduct, err := NewProduct("1", "foo", productPrice)
	assert.NoError(t, err)

	orderAddress, err := NewAddress("test", "test", "test", "test", "test")
	assert.NoError(t, err)

	return orderProduct, orderAddress
}
