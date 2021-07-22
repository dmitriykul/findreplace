package main

import (
	"errors"
	"findreplace/pkg/findreplace/app"
	"findreplace/pkg/findreplace/infrastructure"
	"log"
	"os"
)

func main() {
	err := findReplace(os.Args[1:])
	if err != nil {
		log.Fatalf("fatal error: %v\n", err)
	}
}

func findReplace(args []string) error {
	var findReplacer app.FindReplacer
	switch args[0] {
	case "find":
		var params app.FindParams
		switch len(args) {
		case 2:
			params.Substr = args[1]
			scanner := infrastructure.NewConsoleScannerFactory()
			reporter := infrastructure.NewReporter()
			return findReplacer.FindSubstr(params, scanner, reporter)
		case 3:
			params.Substr = args[1]
			params.Path = args[2]
			scanner:= infrastructure.NewFileScannerFactory()
			reporter := infrastructure.NewReporter()
			return findReplacer.FindSubstr(params, scanner, reporter)
		default:
			return errors.New("missing find arguments: <substr> [<path>]")
		}
	case "replace":
		var params app.ReplaceParams
		switch len(args) {
		case 3:
			params.Substr = args[1]
			params.Replacement = args[2]
			storer := infrastructure.NewFileTextStore()
			reporter := infrastructure.NewReporter()
			scanner := infrastructure.NewConsoleScannerFactory()
			return findReplacer.ReplaceSubstr(params, storer, reporter, scanner)
		case 4:
			params.Substr = args[1]
			params.Replacement = args[2]
			params.Path = args[3]
			storer := infrastructure.NewFileTextStore()
			reporter := infrastructure.NewReporter()
			scanner := infrastructure.NewFileScannerFactory()
			return findReplacer.ReplaceSubstr(params, storer, reporter, scanner)
		default:
			return errors.New("missing replace arguments: <substr> <newStr> [<path>]")
		}
	}
	return nil
}
