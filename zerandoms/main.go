package main

import (
	"fmt"
	"os"
	"strconv"
	zerandom "zerandoms/pkg/random"
)

func main() {

	args := os.Args[1:]
	i, _ := strconv.Atoi(args[0])
	s, _ := zerandom.Random(i)

	fmt.Println(s)

}
