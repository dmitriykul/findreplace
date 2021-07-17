package infrastructure

import (
	"bufio"
	"findreplace/pkg/findreplace/app"
	"os"
)

type consoleScanner struct {
	text string
}

func NewConsoleLineScanner() app.LineScanner {
	return &consoleScanner{}
}

func(l *consoleScanner) ReadLine() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l.text = scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return l.text, nil
}

func(l *consoleScanner) ReadFileLine(path string) error {

	return nil
}
