package main

import (
	"os"
	"fmt"
	"bufio"
//	"strings"
)

func runScript() {
	srcB, err := os.ReadFile(args[0])
	if err != nil { os.Stderr.WriteString("can't read file\n") ; os.Exit(1) }
	run(string(srcB))
	if hadErr { os.Exit(1) }

}
func runPrompt() {
	for {
		p := "\033[38;2;2;240;164m?\033[0m"
		if hadErr { p = "\033[38;2;255;117;127m!\033[0m" }
		fmt.Fprintf(os.Stdout, "\033[38;2;255;215;95m(\033[0m%s\033[38;2;255;215;95m):\033[0m ", p)
		hadErr = false


		reader := bufio.NewReader(os.Stdin)
		li, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "err redaing line: %v", err)
			os.Exit(1)
		}

		if li == "exit()" { break }
		run(li) ; fmt.Print("\n")
	}
}

func run(src string) {
	//what the hell was the point of writing this?
/*	toks := strings.FieldsFunc(src, func(r rune) bool {
		switch string(r) {
     case " ", "\n", "\t": return true
		}; return false
	})*/

	toks := lexToks(src)
	for _, t := range toks { fmt.Print(t.typ.String()+"\n") }
}
