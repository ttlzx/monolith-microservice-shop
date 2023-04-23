package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttlzx/monolith-microservice-shop/internal/common/price"
)

func TestNewProduct(t *testing.T) {
	testPrice, err := price.NewPrice(42, "USD")
	assert.NoError(t, err)

	testCases := []struct {
		TestName string

		ID          ID
		Name        string
		Description string
		Price       price.Price

		ExpectedErr error
	}{
		{
			TestName:    "valid",
			ID:          "1",
			Name:        "foo",
			Description: "bar",
			Price:       testPrice,
		},
		{
			TestName:    "empty_id",
			ID:          "",
			Name:        "foo",
			Description: "bar",
			Price:       testPrice,

			ExpectedErr: ErrEmptyID,
		},
		{
			TestName:    "empty_name",
			ID:          "1",
			Name:        "",
			Description: "bar",
			Price:       testPrice,

			ExpectedErr: ErrEmptyName,
		},
	}

	for _, c := range testCases {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := NewProduct(c.ID, c.Name, c.Description, c.Price)
			assert.EqualValues(t, c.ExpectedErr, err)
		})
	}
}
