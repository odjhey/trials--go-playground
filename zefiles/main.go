package main

import (
	"fmt"
	"zefiles/pkgs/zefiles"
)

func main() {
	fmt.Println("main start")

	const testfile = "testfile.txt"

	rwer, _ := zefiles.New(testfile)

	rwer.Read()
	rwer.Write("Hello another")

}
