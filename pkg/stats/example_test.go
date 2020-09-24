package stats

import (
	"fmt"
	"github.com/bdaler/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{1, 10, "T", types.StatusOk},
		{1, 11, "U", types.StatusFail},
		{1, 12, "R", types.StatusInProgress},
	}
	fmt.Println(Avg(payments))
	//Output:
	// 11
}

func ExampleTotalInCategory_empty() {
	paymets := []types.Payment{}
	fmt.Println(TotalInCategory(paymets, "U"))
	//Output:
	// 0
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{1, 10, "T", types.StatusOk},
		{1, 11, "U", types.StatusFail},
		{1, 12, "R", types.StatusInProgress},
	}
	fmt.Println(TotalInCategory(payments, "U"))
	//Output:
	// 0
}
