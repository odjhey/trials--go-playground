package zefiles

import (
	"fmt"
	"os"
)

/*
q5- Create a package that provide functionality to read, write and close a file. Please store the file connection in a struct.
    Create three methods on the struct (read, write and close) and use this in a main package.
*/

type zefile struct {
	filename string
}

func Test() {
	fmt.Println("Test")
}

func New(filename string) (*zefile, error) {
	return &zefile{filename: filename}, nil
}

func (z zefile) Read() error {
	sometype, err := os.ReadFile(z.filename)
	if err != nil {
		return err
	}

	fmt.Println(string(sometype))

	return nil
}
