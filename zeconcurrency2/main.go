package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	fmt.Println("main")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go run(i)
	}

	wg.Wait()

}

func run(id int) {
	defer wg.Done()
	fmt.Println("id: ", id)
}
