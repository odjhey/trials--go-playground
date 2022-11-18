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

	// call async
	promiseLikeChannel3 := async.Call(func() interface{} {
		return "Hello World! the third"
	})

	// call async
	promiseLikeChannel2 := async.Call(func() interface{} {
		return "Hello World! the second"
	})

	var messages = []interface{}{}

	// simulate await with below construct
	// message := <-promiseLikeChannel

	// call await decending
	messages = append(messages, <-promiseLikeChannel3)
	messages = append(messages, <-promiseLikeChannel2)
	messages = append(messages, <-promiseLikeChannel)

	for _, v := range messages {
		fmt.Println(v)
	}
}
