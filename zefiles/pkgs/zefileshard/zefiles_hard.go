package zefileshard

import (
	"fmt"
	"log"
	"os"
)

/*
q5- Create a package that provide functionality to read, write and close a file. Please store the file connection in a struct.
    Create three methods on the struct (read, write and close) and use this in a main package.
*/

type zefileHard struct {
	filename string
	file     *os.File
}

func New(filename string) (*zefileHard, error) {
	return &zefileHard{filename: filename}, nil
}

func (z *zefileHard) Read() ([]byte, error) {

	var err error
	z.file, err = os.Open(z.filename)
	if err != nil {
		return nil, err
	}

	info, err := z.file.Stat()
	if err != nil {
		return nil, err
	}

	fmt.Println("asdlkfj", info.Size())
	byteSlice := make([]byte, info.Size())
	_, err = z.file.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	log.Printf("Data read: %s\n", byteSlice)

	return byteSlice, nil
}

func (z *zefileHard) Close() error {
	err := z.file.Close()
	return err
}
