package infrastructure

import "findreplace/pkg/findreplace/app"

type consoleScannerFactory struct {

}

func NewConsoleScannerFactory() app.LineScannerFactory {
	return &consoleScannerFactory{}
}

func (c *consoleScannerFactory) CreateScanner(path string) (app.LineScanner, error) {
	return NewConsoleLineScanner(), nil
}
