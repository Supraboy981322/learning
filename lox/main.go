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
  LEFT_PAREN TokType = iota //`(`
	RIGHT_PAREN //`)`
	LEFT_BRACE  //`{`
	RIGHT_BRACE //`}`
  COMMA //`,`
	DOT   //`.`
	MINUS //`-`
	PLUS  //`+`
	SEMICOLON //`;`
	SLASH //`/`
	STAR  //`*`
  BANG  //`!`
	BANG_EQUAL //`!=`
  EQUAL //`=`
	EQUAL_EQUAL //`==`
  GREATER //`>`
	GREATER_EQUAL //`>=`
  LESS  //`<`
	LESS_EQUAL //`<=`
  IDENT //identifier
	STRING //`"..."` or `'...'`
	NUMBER //`1` or `2` or `3` or `4` or `5` or `6` or `7` or `8` or `9` or `0`
  AND   //`&`
	CLASS //`class`
	ELSE  //`:`
	FALSE //`false`
	FN  //`fn`
	FOR //`for`
	TERN //`?` if this affends you, GOOD (I can't understand you... creatures)
	NIL //`nil`
	OR  //`|`
  PRINT //`print`
	RETURN //`return`
	SUPER //`super`
	THIS //`this`
	TRUE //`true`
	VAR  //`var`
	WHILE //`while`
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
