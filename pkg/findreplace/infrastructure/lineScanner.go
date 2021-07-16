package infrastructure

import (
	"bufio"
	"findreplace/pkg/findreplace/app"
	"io"
)

type lineScanner struct {
	text string
}

func NewLineScanner() app.LineScanner {
	return &lineScanner{}
}

func(l *lineScanner) ReadConsoleLine(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	l.text = scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return l.text, nil
}

func(l *lineScanner) ReadFileLine(path string) error {

	return nil
}
