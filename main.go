package main

import (
	"cat/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
