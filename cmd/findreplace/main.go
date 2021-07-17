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
			var scanner app.LineScanner = infrastructure.NewConsoleLineScanner()
			return findReplacer.FindSubstr(params, scanner)
		case 3:
			params.Substr = args[1]
			params.Path = args[2]
			var scanner app.LineScanner = infrastructure.NewConsoleLineScanner()
			return findReplacer.FindSubstr(params, scanner)
		default:
			return errors.New("missing find arguments: <substr> [<path>]")
		}
	case "replace":
		var params app.ReplaceParams
		switch len(args) {
		case 3:
			params.Substr = args[1]
			params.Replacement = args[2]
			return findReplacer.ReplaceSubstr(params)
		case 4:
			params.Substr = args[1]
			params.Replacement = args[2]
			params.Path = args[3]
			return findReplacer.ReplaceSubstr(params)
		default:
			return errors.New("missing replace arguments: <substr> <newStr> [<path>]")
		}
	}
	return nil
}
