package app

import "io"

type Reporter interface {
	PrintLine(str string) error
}

type LineScanner interface {
	ReadConsoleLine(r io.Reader) (string, error)
	ReadFileLine(path string) error
}

type TextStore interface {
	StoreText(text, file string) error
}