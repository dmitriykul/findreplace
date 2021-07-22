package infrastructure

import "findreplace/pkg/findreplace/app"

type fileScannerFactory struct {

}

func NewFileScannerFactory() app.LineScannerFactory {
	return &fileScannerFactory{}
}

func (l *fileScannerFactory) CreateScanner(path string) (app.LineScanner, error) {
	return NewFileScanner(path)
}
