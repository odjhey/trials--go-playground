package main

import (
	"fmt"
	"os"
	"strconv"
	zerandom "zerandoms/pkg/random"
)

func main() {

	args := os.Args[1:]
	seed, _ := strconv.Atoi(args[0])
	freq, _ := strconv.Atoi(args[1])

	for i := 0; i < freq; i++ {
		s, _ := zerandom.Random(seed + i)
		fmt.Println(i+1, "-> ", s)
	}

}
