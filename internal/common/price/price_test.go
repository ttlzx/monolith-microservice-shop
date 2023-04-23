package price

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			ExpectedErr: ErrPriceTooLow,
		},
		{
			Name:        "empty_currency",
			Cents:       10,
			Currency:    "",
			ExpectedErr: ErrInvalidCurrency,
		},
		{
			Name:        "invalid_currency_length",
			Cents:       10,
			Currency:    "US",
			ExpectedErr: ErrInvalidCurrency,
		},
	}

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			_, err := NewPrice(c.Cents, c.Currency)
			assert.EqualValues(t, c.ExpectedErr, err)
		})
	}
}
