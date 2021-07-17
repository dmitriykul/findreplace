package app

type Reporter interface {
	PrintLine(str string) error
}

type LineScanner interface {
	ReadLine() (string, error)
}

type TextStore interface {
	StoreText(text, file string) error
}