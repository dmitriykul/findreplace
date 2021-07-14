package main

import (
	"findreplace/pkg/findreplace/infrastructure"
	"os"
)

func main() {
	infrastructure.Deliver(os.Args[1:])
}