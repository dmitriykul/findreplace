package infrastructure

import (
	"findreplace/pkg/findreplace/app"
	"io"
)

type lineScanner struct {

}

func NewLineScanner() app.LineScanner {
	return &lineScanner{}
}

func(l *lineScanner) ReadConsoleLine(r io.Reader) error {

	return nil
}

func(l *lineScanner) ReadFileLine(path string) error {

}
