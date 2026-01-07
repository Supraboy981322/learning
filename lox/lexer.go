package main

import ("fmt")

type (
	tokLexr struct {
		start int
		line int
		src string
		pos int
		cur string
		toks []Tok
	}
)

var (
)

func lexToks(src string) []Tok {
	p := tokLexr{
		src: src,
		toks: []Tok{},
	}
	return p.lexTok()
}

func (p *tokLexr) lexTok() []Tok {
	if p.eof() { return p.toks }
	p.start = p.pos
	p.cur = string(p.src[p.pos])
	switch p.cur {
	 case ")":  p.addTok(RIGHT_PAREN)
	 case "{":  p.addTok(LEFT_BRACE)
	 case "}":  p.addTok(RIGHT_BRACE)
	 case ",":  p.addTok(COMMA)
	 case ".":  p.addTok(DOT)
	 case "+":  p.addTok(PLUS)
	 case ";":  p.addTok(SEMICOLON)
   case ":":  p.addTok(ELSE)
	 case "?":  p.addTok(TERN)
	 case "&":  p.addTok(AND)
	 case "|":  p.addTok(OR)
	 case "/":  p.addTok(SLASH)
	 case "*":  p.addTok(STAR)
	 case "\"": p.string()
	 case "\n": p.line++
	 case " ", "\r", "\t": //do nothing
	 case "-":
		if p.match("-") { 
			for { if p.match("\n") { break } ; if !p.eof() { p.pos++ } }
		} else { p.addTok(MINUS) }
	 case "!":
		if p.match("=") { p.addTok(BANG_EQUAL) } else { p.addTok(BANG) }
	 case "=":
		if p.match("=") { p.addTok(EQUAL_EQUAL) } else { p.addTok(EQUAL) }
	 case "<":
		if p.match("=") { p.addTok(LESS_EQUAL) } else { p.addTok(LESS) }
	 case ">":
		if p.match("=") { p.addTok(GREATER_EQUAL) } else { p.addTok(GREATER) }
	 case "(":
		if p.match("*") {
			for {
				if p.match("*") && p.match(")") { return p.lexTok() }
				if !p.eof() { if p.src[p.pos] == '\n' { p.line++ } ; p.pos++ }
			}
		} else { p.addTok(LEFT_PAREN) }
	 default:
		if !p.isDig() && !p.isAlpha() {
			eror(p.line, fmt.Sprintf("unexpected char: %s", p.cur))
		} else if p.isAlpha() { p.ident()
		} else { p.num() }
	}
	p.pos++
	return p.lexTok()
}

func (p *tokLexr) eof() bool { return p.pos >= len(p.src) }

func (p *tokLexr) addTok(t TokType) {
	txt := p.src[p.start:p.pos]
	p.toks = append(p.toks, Tok{
		typ: t,
		lexeme: t.String(),
		literal: txt,
		line: p.line,
	})
}

func (p *tokLexr) match(c string) bool {
	if p.eof() || p.pos+1 >= len(p.src) { return false }
	if string(p.src[p.pos+1]) != c { return false }

	p.pos++
	return true
}

func (p *tokLexr) peek() string {
	if p.eof() || p.pos+1 >= len(p.src) { return "" }
	return string(p.src[p.pos+1])
}

func (p *tokLexr) string() {
	for {
		if p.eof() || p.peek() == "\"" { break }
		if p.peek() == "\n" { p.line++ }
		p.cur = string(p.src[p.pos])
		if p.cur == "\\" && p.match("n") {
			p.src = p.src[:p.pos] + "\n" + p.src[p.pos:]
		}
		p.pos++
	}
	if p.eof() {
		eror(p.line, "unterminated string")
		return
	}
	p.start++ //removes begining token
	p.pos++ //corrects for missing last char
	p.addTok(STRING)
}

func (p *tokLexr) isDig() bool {
	switch p.cur {
	 case "0","9","8","7","6","5","4","3","2","1": return true
	}
	return false
}

func (p *tokLexr) num() {
	for {
		if !p.isDig() || p.eof() { break } else {
			p.pos++ ; if p.eof() { break }
			p.cur = string(p.src[p.pos])
		}
	}

	next := p.peek() ; p.pos++
	p.cur = string(p.src[p.pos])
	if next == "." && p.isDig() {
		if p.eof() { return }
		p.cur = string(p.src[p.pos])
		for {
			if !p.isDig() || p.eof() { break } else {
				p.pos++ ; if p.eof() { break }
				p.cur = string(p.src[p.pos])
			}
		}
	}

	p.addTok(NUMBER)
}

func (p *tokLexr) ident() {
	for {
		if p.isAlphaNumeric() {
			p.pos++ ; if p.eof() { break }
			p.cur = string(p.src[p.pos])
		} else { break }
	}

	p.addTok(IDENT)
}

func (p *tokLexr) isAlpha() bool {
	var c rune //since one char, just get first rune 
	for _, r := range p.cur { c = r ; break } 

	return (c >= 'a' && c <= 'z') ||
	       (c >= 'A' && c <= 'Z') ||
	       c == '_'
}

func (p *tokLexr) isAlphaNumeric() bool {
	return p.isAlpha() || p.isDig()
}

func (p *tokLexr) keyword(str string) string {
	words := map[string]TokType {
		"class": CLASS,
		"false": FALSE,
		"fo": FOR,
		"fn": FN,
		"nil": NIL,
		"print": PRINT,
		"return": RETURN,
		"esc": ESC,
		"super": SUPER,
		"this": THIS,
		"true": TRUE,
		"var": VAR,
		"while": WHILE,
	}
}
