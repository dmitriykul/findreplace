package app

type Reporter interface {
	PrintLine(str string) error
}

type LineScanner interface {
	ReadLine() (bool, string, error)
	NewScanner(path string) error
}

type LineScannerChanger interface {
	ChangeScanner(path string) (LineScanner, error)
}

type TextStore interface {
	StoreText(text []byte, file string) error
}