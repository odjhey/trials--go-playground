package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main")

	c := make(chan int)

	go sendVal(c, 12)
	go sendVal(c, 14)
	go sendVal(c, 3)

	sum := sum(c)
	fmt.Println("sum", sum)

}

func sendVal(c chan int, v int) {
	time.Sleep(2 * time.Second)
	c <- v
}

func sum(c chan int) int {

	v1, v2, v3 := <-c, <-c, <-c

	return v1 + v2 + v3

}
