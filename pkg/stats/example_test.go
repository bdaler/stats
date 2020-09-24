package stats

import (
	"fmt"
	"github.com/bdaler/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{1, 10, "T"},
		{1, 11, "U"},
		{1, 12, "R"},
	}
	fmt.Println(Avg(payments))
	//Output:
	// 11
}
