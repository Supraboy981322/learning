package main

import (
	"os"
)

type (
	TokType int
	Tok struct {
		typ TokType
		lexeme string
		line int
		literal string
	}
)

var (
	hadErr bool
	args = os.Args[1:]
)

func main() {
	if len(args) > 1 {
		os.Stderr.WriteString("too many args, just need script\n")
		os.Exit(1)
	} else if len(args) == 1 { runScript() } else { runPrompt() }
}
