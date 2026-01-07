package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "need dir\n")
		os.Exit(1)
	}; outputDir := os.Args[1]

	if err := defAst(outputDir, "Expr", []string{
		"Assign : name *Token, value Expr",
		"Binary : left Expr, operator *Token, right Expr",
		"Call : callee Expr, paren *Token, arguments []Expr",
		"Get : object Expr, name *Token",
		"Grouping : expression Expr",
		"Literal : value interface{}",
		"Logical : left Expr, operator *Token, right Expr",
		"Set : object Expr, name *Token, value Expr",
		"Super : keyword *Token, method *Token",
		"This : keyword *Token",
		"Unary : operator *Token, right Expr",
		"Variable : name *Token",
	}); err != nil { fmt.Println(err) ; os.Exit(1) }

	if err := defAst(outputDir, "Stmt", []string{
		"Block : statements []Stmt",
		"Class : name *Token, superclass *Variable, methods []*Function",
		"Expression: expression Expr",
		"Function : name *Token, params []*Token, body []Stmt",
		"If : condition Expr, thenBranch Stmt, elseBranch Stmt",
		"Include : path *Token",
		"Print : expression Expr",
		"Return : keyword *Token, value Expr",
		"Var : name *Token, initializer Expr",
		"While : condition Expr, body Stmt",
	}); err != nil { fmt.Println(err) ; os.Exit(1) }
}

func defAst(outputDir string, baseName string, types []string) error {
	var path string ; {
		bN := strings.ToLower(baseName)+".go"
		path = filepath.Join(outputDir, bN)
	}

	//hmmm....
	//
	//	var file; ({ file, err = sys.mkfile(path) } |err| != nil ) ? {
	//		file.close() ; return err;
	//	}; defer file.close();
	//
	//(i like the syntax here)

	var file *os.File ; var err error
	if file, err = os.Create(path); err != nil {
		file.Close() ; return err
	}
	fmt.Println(path)

	wr := bufio.NewWriter(file)
	defer wr.Flush()

	for _, l := range []string{
		"package main\n",
		"type "+baseName+" interface {",
		"  Accept(Visitor"+baseName+") (interface{}, error)",
		"  IsType(interface{}) bool",
		"}\n",
	} {	wr.Write([]byte(l+"\n")) }

	defVis(wr, baseName, types)

	for _, typ := range types {
		t := strings.Split(typ, ":")
		className := strings.TrimSpace(t[0])
		fields := strings.TrimSpace(t[1])

		defTyp(wr, baseName, className, fields)
		defIsTyp(wr, className)
	}

	return nil
}

func defTyp(wr *bufio.Writer, baseName, className, fields string) {
	wr.Write([]byte("type "+className+" struct {\n"))

	fieldList := strings.Split(fields, ", ")
	for _, field := range fieldList {
		vS := strings.Split(field, " ")
		for i, v := range vS {
			if i == 0 {
				wr.Write([]byte(strings.Title(v)+" "))
			} else { wr.Write([]byte(v+"\n"))	}
		}
	}

	for _, l := range []string{ 
		"}\n",
		"func New"+className+"("+fields+") "+baseName+"{",
		"  return &"+className+"{",
	} { wr.Write([]byte(l+"\n")) }

	args := make([]string, 0)
	for _, field := range fieldList {
		name := strings.Split(field, " ")[0]
		args = append(args, name)
	}

	{
		cName := strings.ToUpper(string(className))
		for _, l := range []string{ 
			strings.Join(args, ",")+",",
			"  }",
			"}\n",
			"func ("+cName+" *"+className+") Accept(visitor Visitor"+
            baseName+") (interface{}, error) {",
			"  return visitor.visit"+className+baseName+"("+cName+")",
			"}\n",
		}{ wr.Write([]byte(l+"\n")) }
	}
}

func defVis(wr *bufio.Writer, baseName string, types []string) {
	wr.Write([]byte("type Visitor"+baseName+" interface {\n"))
	{
		bL := []string{"  visit",baseName, "(*", ") (interface{}, error)\n" }
		for _, typ := range types {
			tN := strings.TrimSpace(strings.Split(typ, ":")[0])
			wr.Write([]byte(bL[0] + tN+bL[1] + bL[2] + tN + bL[3]))
		}
	}
	wr.Write([]byte("}\n\n"))
}

func defIsTyp(wr *bufio.Writer, className string) {
	for _, l := range []string {
		"\nfunc (rec *"+className+") IsType(v interface{}) bool {",
		"  switch v.(type) {",
		"   case *"+className+": return true",
		"  }",
		"  return false",
		"}",
	} { wr.Write([]byte(l+"\n")) }
}
