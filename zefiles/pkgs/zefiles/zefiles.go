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

func (z zefile) Read() ([]byte, error) {
	somebytes, err := os.ReadFile(z.filename)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(somebytes))

	return somebytes, nil
}

func (z zefile) Write(text string) error {
	err := os.WriteFile(z.filename, []byte(text), 6)
	return err
}
