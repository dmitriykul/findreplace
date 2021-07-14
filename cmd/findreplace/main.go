package main

import (
	"findreplace/pkg/findreplace/app"
	"os"
)

func main() {
	app.FindReplace(os.Args[1:])
}