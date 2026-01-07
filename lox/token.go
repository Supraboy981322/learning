package main

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
	ESC //`esc` (arbitrarily eccape blocks)
	CLOSE //`CLOSE` (quit program)
	INVALID //used internally
	NOP //no-op (eg: `_ = funcWithReturn()` or `(a > 0) ? { _ } : { print "a is smol" }`)
  EOF

	IF TokType = TERN
)

func (t *Tok) ToString() string {
	return t.typ.String() + " " + t.lexeme + " " + t.literal
}

func (t TokType) String() string {
	switch t {
	 case LEFT_PAREN:   return "LEFT_PAREN" 
	 case RIGHT_PAREN:  return "RIGHT_PAREN" 
	 case LEFT_BRACE:   return "LEFT_BRACE" 
	 case RIGHT_BRACE:  return "RIGHT_BRACE" 
	 case COMMA:        return "COMMA" 
	 case DOT:          return "DOT" 
	 case MINUS:        return "MINUS" 
	 case PLUS:         return "PLUS" 
	 case SEMICOLON:    return "SEMICOLON" 
	 case SLASH:        return "SLASH" 
	 case STAR:         return "STAR" 
	 case BANG:         return "BANG" 
	 case BANG_EQUAL:   return "BANG_EQUAL" 
	 case EQUAL:        return "EQUAL" 
	 case EQUAL_EQUAL:  return "EQUAL_EQUAL" 
	 case GREATER:      return "GREATER" 
	 case GREATER_EQUAL:return "GREATER_EQUAL"
	 case LESS:         return "LESS" 
	 case LESS_EQUAL:   return "LESS_EQUAL" 
	 case IDENT:   			return "IDENT" 
	 case STRING:       return "STRING" 
	 case NUMBER:       return "NUMBER" 
	 case AND:          return "AND" 
	 case CLASS:        return "CLASS" 
	 case ELSE:         return "ELSE" 
	 case FALSE:        return "FALSE" 
	 case FN:           return "FN" 
	 case FOR:          return "FOR" 
	 case TERN:         return "TERN"
	 case NIL:          return "NIL" 
	 case OR:           return "OR" 
	 case PRINT:        return "PRINT" 
	 case RETURN:       return "RETURN" 
	 case SUPER:        return "SUPER" 
	 case THIS:         return "THIS" 
	 case TRUE:         return "TRUE" 
	 case VAR:          return "VAR" 
	 case WHILE:        return "WHILE" 
	 case ESC:          return "ESC"
	 case EOF:          return "EOF" 
	 case CLOSE:         return "CLOSE"
	 case INVALID:      return "INVALID"
	}
	return "unknown token"
}
