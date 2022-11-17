package worker

import "fmt"

func TotalExpenses(emps ...MoneySpender) int {
	var sum = 0
	for _, e := range emps {
		sum = sum + e.GetMonthlyExpenses()
	}
	return sum
}

func PrintExpenses(emps ...MoneySpender) {
	for _, e := range emps {
		fmt.Println(e.GetName(), e.GetMonthlyExpenses())
	}
}
