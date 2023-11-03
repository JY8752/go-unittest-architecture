//go:build unit

package domain_test

import (
	"testing"

	"github.com/JY8752/go-unittest-architecture/domain"
	"github.com/stretchr/testify/assert"
)

func TestDraw(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		Weights []int
		Expect  domain.ItemId
	}{
		"case1: when seed 10 total weights 16, rand 14": {
			Weights: []int{1, 5, 10},
			Expect:  domain.NewItemId(3),
		},
		"case2: when seed 10 total weights 16, rand 14": {
			Weights: []int{1, 10, 5},
			Expect:  domain.NewItemId(3),
		},
		"case3: when seed 10 total weights 16, rand 14": {
			Weights: []int{5, 1, 10},
			Expect:  domain.NewItemId(3),
		},
		"case4: when seed 10 total weights 16, rand 14": {
			Weights: []int{5, 10, 1},
			Expect:  domain.NewItemId(2),
		},
		"case5: when seed 10 total weights 16, rand 14": {
			Weights: []int{10, 1, 5},
			Expect:  domain.NewItemId(3),
		},
		"case6: when seed 10 total weights 16, rand 14": {
			Weights: []int{10, 5, 1},
			Expect:  domain.NewItemId(2),
		},
		"case1: only one item": {
			Weights: []int{1},
			Expect:  domain.NewItemId(1),
		},
		"case2: only one item": {
			Weights: []int{1000},
			Expect:  domain.NewItemId(1),
		},
	}

	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			// Arange
			t.Parallel()
			gacha := domain.NewGacha(newGachaItemWeights(t, tt.Weights))
			seed := int64(10)

			// Act
			itemId, err := gacha.Draw(seed)
			if err != nil {
				t.Fatal(err)
			}

			// Asertion
			assert.Equal(t, tt.Expect, itemId)
		})
	}
}

func newGachaItemWeights(t *testing.T, weights []int) domain.GachaItemWeights {
	t.Helper()
	gachaItemWeights := make(domain.GachaItemWeights, len(weights))
	for i, w := range weights {
		gachaItemWeights[i] = struct {
			ItemId domain.ItemId
			Weight int
		}{
			ItemId: domain.NewItemId(int64(i + 1)),
			Weight: w,
		}
	}
	return gachaItemWeights
}
