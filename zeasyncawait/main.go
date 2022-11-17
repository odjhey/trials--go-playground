package main

import (
	"fmt"
	"zeasyncawait/async"
)

func main() {
	fmt.Println("main")

	// call async
	promiseLikeChannel := async.Call(func() interface{} {
		return "Hello World!"
	})

	// simulate await
	message := <-promiseLikeChannel

	fmt.Println(message)
}
