package utils

import "fmt"

func FormatCurrency(amount int) string {
	formattedAmount := fmt.Sprintf("%d", amount)
	for i := len(formattedAmount) - 3; i > 0; i -= 3 {
		formattedAmount = formattedAmount[:i] + "." + formattedAmount[i:]
	}
	return formattedAmount
}
