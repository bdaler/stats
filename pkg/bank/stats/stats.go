package stats

import "bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	count := len(payments)
	sum := 0

	for i := 0; i < count; i++ {
		sum += int(payments[i].Amount)
	}
	avg := sum / count

	return types.Money(avg)
}
