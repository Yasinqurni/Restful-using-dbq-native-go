package entity_test

import (
	"restful-api/domain/entity"
	"restful-api/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemEntity(t *testing.T) {
	type expected struct {
		Name     string
		Category string
		Price    uint64
		Quantity uint32
	}
	type testCases struct {
		name     string
		dto      entity.ItemDTO
		expected *expected
		isError  bool
		err      any
	}
	testcases := []testCases{
		{
			name: "positive",
			dto: entity.ItemDTO{
				Name:     "realme 5 pro",
				Category: "handphone",
				Price:    2500000,
				Quantity: 10,
			},
			expected: &expected{
				Name:     "realme 5 pro",
				Category: "handphone",
				Price:    2500000,
				Quantity: 10,
			},
			isError: false,
			err:     nil,
		},
		{
			name: "negative",
			dto: entity.ItemDTO{
				Name:     "",
				Category: "handphone",
				Price:    2500000,
				Quantity: 10,
			},
			expected: nil,
			isError:  true,
			err:      pkg.EMPTY_NAME,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {

			item, err := entity.NewItem(test.dto)

			if !test.isError {
				assert.Equal(t, test.expected.Name, item.GetName())
				assert.Equal(t, test.expected.Category, item.GetCategory())
				assert.Equal(t, test.expected.Price, item.GetPrice())
				assert.Equal(t, test.expected.Quantity, item.GetQty())
			} else {
				assert.Equal(t, test.err, err.Error())
			}

		})
	}
}
