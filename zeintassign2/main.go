package main

import (
	"fmt"
	"zeintassign2/worker"
)

func main() {
	factory := worker.WorkerFactory{}

	perm, _ := factory.NewHire("John", 9400)
	contrator, _ := factory.NewContract("Syna", 12000)
	intern, _ := factory.NewIntern("Pho", 1800)

	worker.PrintExpenses(perm, contrator, intern)
	fmt.Println()
	totalExpenses := worker.TotalExpenses(perm, contrator, intern)

	fmt.Println("Total expenses: ", totalExpenses)
}
