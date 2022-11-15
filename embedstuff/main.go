package main

import "fmt"

type user struct {
	name string
}

type admin struct {
	user
	adminStuff string
}

func (u *user) ChangeName(newname string) {
	u.name = newname
}

func main() {
	fmt.Println("1", user{name: "john"})
	a := admin{user: user{name: "asdkfj"}, adminStuff: "xxx"}
	fmt.Println("2", a)
	fmt.Println("3", a.user)

	accept(a.user)

	// ChangeName is implemented by user
	a.ChangeName("noice! new name")

	fmt.Println("Is works? " + a.name)
}

func accept(u user) {
	fmt.Println("zz", u)
}
