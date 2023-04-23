package price_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttlzx/monolith-microservice-shop/internal/common/price"
)

func TestNewPrice(t *testing.T) {
	testCases := []struct {
		Name        string
		Cents       uint
		Currency    string
		ExpectedErr error
	}{
		{
			Name:     "valid",
			Cents:    10,
			Currency: "EUR",
		},
		{
			Name:        "invalid_cents",
			Cents:       0,
			Currency:    "EUR",
			ExpectedErr: price.ErrPriceTooLow,
		},
		{
			Name:        "empty_currency",
			Cents:       10,
			Currency:    "",
			ExpectedErr: price.ErrInvalidCurrency,
		},
		{
			Name:        "invalid_currency_length",
			Cents:       10,
			Currency:    "US",
			ExpectedErr: price.ErrInvalidCurrency,
		},
	}

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			_, err := price.NewPrice(c.Cents, c.Currency)
			assert.EqualValues(t, c.ExpectedErr, err)
		})
	}
}
