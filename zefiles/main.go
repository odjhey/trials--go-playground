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

	hfile, err := zefileshard.New(testfile)
	if err != nil {
		fmt.Println("Error")
		return
	}
	defer hfile.Close()

	rbytes, err := hfile.Read()
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Printf("%s\n", rbytes)

	err = hfile.Write("John Cena!")
	if err != nil {
		fmt.Println("Write Error")
		return
	}

	// eof was read, q: how to put back the cursor to the beginning?
	// confirmed, use the `readAt` method daw
	rbytes2, err := hfile.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", rbytes2)

}
