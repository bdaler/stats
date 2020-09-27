package stats

import (
	"github.com/bdaler/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "it", Amount: 1_000_00},
		{ID: 2, Category: "it", Amount: 1_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "food", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"it":   100_000,
		"auto": 350_000,
		"food": 500_000,
	}

	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
