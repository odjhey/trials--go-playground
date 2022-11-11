package main

import (
	"fmt"
	"theslicephase/pkg/inspect"
)

func main() {

	i := []int{1, 2, 3}

	inspect.InspectSlice("i", i)
	b := i[1:2]
	inspect.InspectSlice("b", b)

	fmt.Println(100_000)

}
