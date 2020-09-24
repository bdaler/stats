package stats

import (
	"github.com/bdaler/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	count := len(payments)
	sum := 0

	for i := 0; i < count; i++ {
		if payments[i].Status != types.StatusFail {
			sum += int(payments[i].Amount)
		}
	}
	avg := sum / count

	return types.Money(avg)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}
	return sum
}
