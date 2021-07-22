package app

type Reporter interface {
	PrintLine(str string) error
}

type LineScanner interface {
	ReadLine() (bool, string, error)
	GetFileName() string
}

type TextStore interface {
	StoreText(text []byte, file string) error
}

type LineScannerFactory interface {
	CreateScanner(path string) (LineScanner, error)
}