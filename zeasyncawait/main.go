package main

import "fmt"

func main() {
	fmt.Println("main")

	// call async
	promiseLikeChannel := callAsync(func() interface{} {
		return "Hello World!"
	})

	// simulate await
	message := <-promiseLikeChannel

	fmt.Println(message)
}

type asyncFn func() interface{}

func callAsync(fn asyncFn) chan interface{} {

	c := make(chan interface{})

	go func() {
		c <- fn()
	}()

	return c

}
