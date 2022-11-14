package main

import "fmt"

type user struct {
	name string
}

type admin struct {
	user
	adminStuff string
}

func main() {
	fmt.Println("1", user{name: "john"})
	a := admin{user: user{name: "asdkfj"}, adminStuff: "xxx"}
	fmt.Println("2", a)
	fmt.Println("3", a.user)

	accept(a.user)
}

func accept(u user) {
	fmt.Println("zz", u)
}
