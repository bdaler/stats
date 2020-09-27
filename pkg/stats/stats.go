package stats

import (
	"github.com/bdaler/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	count := 0
	sum := 0

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		sum += int(payment.Amount)
		count++
	}
	avg := sum / (count)

	return types.Money(avg)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	result := make(map[types.Category]types.Money)
	for _, payment := range payments {
		if _, err := result[payment.Category]; err {
			continue
		}
		filtered := PaymentByCategory(payments, payment.Category)
		result[payment.Category] = Avg(filtered)
	}
	return result
}

func PaymentByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}
