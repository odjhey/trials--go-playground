package main

import (
	"fmt"
	"zefiles/pkgs/zefileshard"
)

func main() {
	fmt.Println("main start")

	const testfile = "testfile.txt"

	/*
		rwer, _ := zefiles.New(testfile)

		rwer.Read()
		rwer.Write("Hello another")
	*/

	hfile, _ := zefileshard.New(testfile)

	rbytes, err := hfile.Read()
	if err != nil {
		fmt.Println("Error")
	}
	defer hfile.Close()

	fmt.Printf("%s", rbytes)

}
