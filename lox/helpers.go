package main

import ("fmt";"os")

func eror(p *TokLexr, msg string) {
	li := "\t"+p.srcLis[p.line]
	report(p.line, li, msg)
}

func report(line int, where string, msg string) {
	//this could've been a ternary.
	//  lang designers, what strange creatures
	var w string 
	if where == "" { w = "" } else { w = "\n" }
	fmt.Fprintf(os.Stderr, "[line %d] err: %s%s%s\n", line, msg, w, where)

	hadErr = true
}
