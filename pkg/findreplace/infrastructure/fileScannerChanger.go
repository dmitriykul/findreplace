package infrastructure

import (
	"bufio"
	"findreplace/pkg/findreplace/app"
	"os"
)

type fileScannerChanger struct {

}

func NewFileScannerChanger() app.LineScannerChanger {
	return &fileScannerChanger{}
}

func (f *fileScannerChanger) ChangeScanner(path string) (app.LineScanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	return &fileScanner{scanner: scanner, file: file, res: true}, nil
}
