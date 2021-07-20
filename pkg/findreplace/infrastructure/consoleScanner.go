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

func(l *consoleScanner) ReadLine() (bool, string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	res := scanner.Scan()
	l.text = scanner.Text()
	if err := scanner.Err(); err != nil {
		return res, "", err
	}

	return res, l.text, nil
}

func(l *consoleScanner) NewScanner(path string) error {
	return nil
}