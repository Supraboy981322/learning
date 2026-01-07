package main

import ("fmt";"os")

func eror(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where string, msg string) {
	//this could've been a ternary.
	//  lang designers, what strange creatures
	var w string 
	if where == "" { w = "" } else { w = " " }

	fmt.Fprintf(os.Stderr, "[line %d] err%s%s: %s\n", line, w, where, msg)

	hadErr = true
}
