package infrastructure

import (
	"findreplace/pkg/findreplace/app"
	"fmt"
)

type consoleReporter struct {

}

func NewReporter() app.Reporter {
	return &consoleReporter{}
}

func(r *consoleReporter) PrintLine(str string) error {
	fmt.Println(str)

	return nil
}