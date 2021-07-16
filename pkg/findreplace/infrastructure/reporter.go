package infrastructure

import "findreplace/pkg/findreplace/app"

type reporter struct {

}

func NewReporter() app.Reporter {
	return &reporter{}
}

func(r *reporter) PrintLine(str string) error {


	return nil
}