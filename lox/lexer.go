package main

import ("os";"strings")

type (
	TokLexr struct {
		start int
		line int
		src string
		pos int
		cur string
		toks []Tok
		srcLis []string
	}
)

var (
)

func lexToks(src string) []Tok {
	srcLis := strings.Split(src, "\n")
	p := TokLexr{
		src: src,
		toks: []Tok{},
		srcLis: srcLis,
	}
	return p.lexTok()
}

func (p *TokLexr) lexTok() []Tok {
	if p.eof() { return p.toks }
	p.start = p.pos
	p.cur = string(p.src[p.pos])
	switch p.cur {
	 case "_":  p.addTok(NOP)
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
	 case "-": p.addTok(MINUS)
	 case "\"": p.string()
	 case "\n": p.line++
	 case " ", "\r", "\t": //do nothing
	 case "~": //the equivalent Zig code looks much better for this one
		if p.match("~") { 
			for { if p.match("\n") { break } ; if !p.eof() { p.pos++ } }
		} else { eror(p, "unexpected char: "+p.cur) }
	 case "!": //insert rant about not having a ternary here
		if p.match("=") { p.addTok(BANG_EQUAL) } else { p.addTok(BANG) }
	 case "=": //insert rant about not having a ternary here
		if p.match("=") { p.addTok(EQUAL_EQUAL) } else { p.addTok(EQUAL) }
	 case "<": //insert rant about not having a ternary here
		if p.match("=") { p.addTok(LESS_EQUAL) } else { p.addTok(LESS) }
	 case ">": //insert rant about not having a ternary here
		if p.match("=") { p.addTok(GREATER_EQUAL) } else { p.addTok(GREATER) }
	 case "(": 
		if p.match("*") {
			if p.multi_com(0) > 0 { eror(p, "multi-line never closed, reached EOF") } 
		} else { p.addTok(LEFT_PAREN) }
	 default:
		if !p.isDig() && !p.isAlpha() {
			eror(p, "unexpected char: "+p.cur)
		} else if p.isAlpha() { p.ident()
		} else { p.num() }
	}
	p.pos++
	return p.lexTok()
}

func (p *TokLexr) eof() bool { return p.pos >= len(p.src) }

func (p *TokLexr) addTok(t TokType) {
	txt := p.src[p.start:p.pos]
	p.toks = append(p.toks, Tok{
		typ: t,
		lexeme: t.String(),
		literal: txt,
		line: p.line,
	})
}

func (p *TokLexr) match(c string) bool {
	if p.eof() || p.pos+1 >= len(p.src) { return false }
	if string(p.src[p.pos+1]) != c { return false }

	p.pos++
	return true
}

func (p *TokLexr) peek() string {
	if p.eof() || p.pos+1 >= len(p.src) { return "" }
	return string(p.src[p.pos+1])
}

func (p *TokLexr) string() {
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
		eror(p, "unterminated string")
		return
	}
	p.start++ //removes begining token
	p.pos++ //corrects for missing last char
	p.addTok(STRING)
}

func (p *TokLexr) isDig() bool {
	switch p.cur {
	 case "0","9","8","7","6","5","4","3","2","1": return true
	}
	return false
}

func (p *TokLexr) num() {
	for {
		if !p.isDig() || p.eof() { break } else {
			p.pos++ ; if p.eof() { break }
			p.cur = string(p.src[p.pos])
		}
	}

	next := p.peek()
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

func (p *TokLexr) ident() {
	for {
		if p.isAlphaNumeric() {
			p.pos++ ; if p.eof() { break }
			p.cur = string(p.src[p.pos])
		} else { break }
	}

	txt := p.src[p.start:p.pos]
	typ := p.keyword(txt) 
	if typ == INVALID { typ = IDENT }

	p.addTok(typ)
}

func (p *TokLexr) isAlpha() bool {
	var c rune //since one char, just get first rune 
	for _, r := range p.cur { c = r ; break } 

	return (c >= 'a' && c <= 'z') ||
	       (c >= 'A' && c <= 'Z') ||
	       c == '_'
}

func (p *TokLexr) isAlphaNumeric() bool {
	return p.isAlpha() || p.isDig()
}

func (p *TokLexr) keyword(str string) TokType {
	words := map[string]TokType {
		"class": CLASS,
		"false": FALSE,
		"for": FOR,
		"fn": FN,
		"nil": NIL,
		"print": PRINT,
		"return": RETURN,
		"esc": ESC,
		"super": SUPER,
		"this": THIS,
		"true": TRUE,
		"var": VAR,
		"close": CLOSE,
		"while": WHILE,
	}

	for k, v := range words { 
		if str == k { return v }
	} ; return INVALID
}

func (p *TokLexr) multi_com(numDeep int) int {
	for {
		if p.match("*") && p.match(")") { break }
		if p.match("(") && p.match("*") {
			numDeep++ ; numDeep = p.multi_com(numDeep)
		}
		if numDeep == -1 {
			eror(p, "multi_com(): numDeep == -1") ; os.Exit(1)
		}
		if !p.eof() {
			if p.src[p.pos] == '\n' { p.line++ }
			p.pos++
		} else { return numDeep+1 }
	}
	return numDeep-1
}
