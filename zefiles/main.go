package main

import (
	"fmt"
	"zefiles/pkgs/zefiles"
)

func main() {
	fmt.Println("main start")

	const testfile = "testfile.txt"

	reader, _ := zefiles.New(testfile)

	reader.Read()

}
