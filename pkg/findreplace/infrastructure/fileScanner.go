package infrastructure

import (
	"bufio"
	"findreplace/pkg/findreplace/app"
	"os"
)

type fileScanner struct {
	file    *os.File
	scanner *bufio.Scanner
	res     bool
}

func NewFileScanner(path string) (app.LineScanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	return &fileScanner{scanner: scanner, file: file, res: true}, nil
}

func (f *fileScanner) ReadLine() (bool, string, error) {
	f.res = f.scanner.Scan()
	if f.res != true {
		f.file.Close()
		return f.res, "", nil
	}

	if err := f.scanner.Err(); err != nil {
		return f.res, "", err
	}

	return f.res, f.scanner.Text(), nil
}

func (f *fileScanner) NewScanner(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	f.scanner = bufio.NewScanner(file)

	return nil
}
