package mock

import (
	"findreplace/pkg/findreplace/app"
)

type consoleScanner struct {
	text string
}

func NewConsoleLineScanner() app.LineScanner {
	return &consoleScanner{}
}

func(l *consoleScanner) ReadLine() (bool, string, error) {
	l.text = "papapa"

	return false, l.text, nil
}

func(l *consoleScanner) GetFileName() string {
	return ""
}