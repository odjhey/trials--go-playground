package zefileshard

import (
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

	file, err := os.OpenFile(filename, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	return &zefileHard{filename: filename, file: file}, nil
}

func (z *zefileHard) Read() ([]byte, error) {

	info, err := z.file.Stat()
	if err != nil {
		return nil, err
	}

	byteSlice := make([]byte, info.Size())
	_, err = z.file.Read(byteSlice)
	if err != nil {
		return nil, err
	}

	return byteSlice, nil
}

func (z *zefileHard) Write(input string) error {
	_, err := z.file.Write([]byte(input))

	return err
}

func (z *zefileHard) Close() error {
	err := z.file.Close()
	return err
}
