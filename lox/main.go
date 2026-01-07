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

const (
	// Single-character tokens.
  LEFT_PAREN TokType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
  COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
  BANG
	BANG_EQUAL
  EQUAL
	EQUAL_EQUAL
  GREATER
	GREATER_EQUAL
  LESS
	LESS_EQUAL
  IDENT
	STRING
	NUMBER
  AND
	CLASS
	ELSE
	FALSE
	FN
	FOR
	TERN //if this affends you, GOOD (I can't understand you... creatures)
	NIL
	OR
  PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	COLON
	ESC //arbitrarily eccape blocks
  EOF
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
