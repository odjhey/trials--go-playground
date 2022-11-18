package main

import (
	"fmt"
	"time"
	"zeasyncawait/async"
)

func main() {
	fmt.Println("main")

	// call async

	promiseLikeChannel := async.Call(func() interface{} {
		fmt.Println("Exec 1st")
		return "Hello World!"
	})

	// call async
	promiseLikeChannel3 := async.Call(func() interface{} {
		fmt.Println("Exec 3rd")
		return "Hello World! the third"
	})

	// call async
	promiseLikeChannel2 := async.Call(func() interface{} {
		fmt.Println("Exec 2nd")
		// simulate expensive computation/network call
		time.Sleep(5 * time.Second)

		return "Hello World! the second"
	})

	var messages = []interface{}{}

	// simulate await with below construct
	// message := <-promiseLikeChannel

	// call await decending
	messages = append(messages, <-promiseLikeChannel3)
	fmt.Println("Received: hello 3rd")
	messages = append(messages, <-promiseLikeChannel2)
	fmt.Println("Received: hello 2nd")
	messages = append(messages, <-promiseLikeChannel)
	fmt.Println("Received: hello 1st")

	fmt.Printf("\nmessages: \n")
	for _, v := range messages {
		fmt.Println(v)
	}
}
