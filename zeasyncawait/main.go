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

	// simulate await/unwrap with below construct
	// message := <-promiseLikeChannel

	// call await decending
	messages = append(messages, <-promiseLikeChannel3)
	messages = append(messages, <-promiseLikeChannel2)
	messages = append(messages, <-promiseLikeChannel)

	for _, v := range messages {
		fmt.Println(v)
	}
}

// NOTES:
// promiseLikeChannel is the term used for JS(or other) devs familiar
// with the concept of Promises/Futures/Tasks have a nicer time
// i personally go with naming it with what it is in
// go, a channel, so channel1, channel2,
//   or channelForMessage, or wrappedValue, or valueInChannel
