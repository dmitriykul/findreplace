package infrastructure

import (
	"bufio"
	"findreplace/pkg/findreplace/app"
	"os"
	"path/filepath"
)

type fileScanner struct {
	file    *os.File
	scanner *bufio.Scanner
	res     bool
	fileName string
}

func NewFileScanner(path string) (app.LineScanner, error) {
	fileName := filepath.Base(path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	return &fileScanner{scanner: scanner, file: file, res: true, fileName: fileName}, nil
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

func (f *fileScanner) GetFileName() string {
	return f.fileName
}
