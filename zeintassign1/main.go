package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Print("main\n\n")
	fmt.Println("call go(1)")
	sum, err := do(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("sum is: ", sum, "\n\n")

	fmt.Println("call go(1,2)")
	sum, err = do(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("sum is: ", sum, "\n\n")

	fmt.Println("call go(1,2,3)")
	sum, err = do(1, 2, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("sum is: ", sum, "\n\n")
}

func do(values ...interface{}) (int, error) {

	fmt.Println("args", values, "len", len(values))
	if len(values) != 2 {
		return 0, errors.New("Please supply more 2 args")
	}

	var sum = 0
	for _, v := range values {
		switch t := v.(type) {
		case int:
			sum = sum + t
		default:
			return 0, errors.New("Invalid types")
		}
	}

	return sum, nil

}
