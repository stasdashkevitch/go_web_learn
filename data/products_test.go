package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	t.Run("valid struct", func(t *testing.T) {
		p := Product{
			Name:  "coffe",
			Price: 1.0,
			SKU:   "aa-aaa-aa",
		}

		err := p.Validate()

		if err != nil {
			t.Fatal(err)
		}
	})
}
