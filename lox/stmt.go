package main

type Stmt interface {
  Accept(VisitorStmt) (interface{}, error)
  IsType(interface{}) bool
}

type VisitorStmt interface {
  visitBlockStmt(*Block) (interface{}, error)
  visitClassStmt(*Class) (interface{}, error)
  visitExpressionStmt(*Expression) (interface{}, error)
  visitFunctionStmt(*Function) (interface{}, error)
  visitIfStmt(*If) (interface{}, error)
  visitIncludeStmt(*Include) (interface{}, error)
  visitPrintStmt(*Print) (interface{}, error)
  visitReturnStmt(*Return) (interface{}, error)
  visitVarStmt(*Var) (interface{}, error)
  visitWhileStmt(*While) (interface{}, error)
}

type Block struct {
Statements []Stmt
}

func NewBlock(statements []Stmt) Stmt{
  return &Block{
statements,
  }
}

func (BLOCK *Block) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitBlockStmt(BLOCK)
}


func (rec *Block) IsType(v interface{}) bool {
  switch v.(type) {
   case *Block: return true
  }
  return false
}
type Class struct {
Name *Tok
Superclass *Variable
Methods []*Function
}

func NewClass(name *Tok, superclass *Variable, methods []*Function) Stmt{
  return &Class{
name,superclass,methods,
  }
}

func (CLASS *Class) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitClassStmt(CLASS)
}


func (rec *Class) IsType(v interface{}) bool {
  switch v.(type) {
   case *Class: return true
  }
  return false
}
type Expression struct {
Expression Expr
}

func NewExpression(expression Expr) Stmt{
  return &Expression{
expression,
  }
}

func (EXPRESSION *Expression) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitExpressionStmt(EXPRESSION)
}


func (rec *Expression) IsType(v interface{}) bool {
  switch v.(type) {
   case *Expression: return true
  }
  return false
}
type Function struct {
Name *Tok
Params []*Tok
Body []Stmt
}

func NewFunction(name *Tok, params []*Tok, body []Stmt) Stmt{
  return &Function{
name,params,body,
  }
}

func (FUNCTION *Function) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitFunctionStmt(FUNCTION)
}


func (rec *Function) IsType(v interface{}) bool {
  switch v.(type) {
   case *Function: return true
  }
  return false
}
type If struct {
Condition Expr
ThenBranch Stmt
ElseBranch Stmt
}

func NewIf(condition Expr, thenBranch Stmt, elseBranch Stmt) Stmt{
  return &If{
condition,thenBranch,elseBranch,
  }
}

func (IF *If) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitIfStmt(IF)
}


func (rec *If) IsType(v interface{}) bool {
  switch v.(type) {
   case *If: return true
  }
  return false
}
type Include struct {
Path *Tok
}

func NewInclude(path *Tok) Stmt{
  return &Include{
path,
  }
}

func (INCLUDE *Include) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitIncludeStmt(INCLUDE)
}


func (rec *Include) IsType(v interface{}) bool {
  switch v.(type) {
   case *Include: return true
  }
  return false
}
type Print struct {
Expression Expr
}

func NewPrint(expression Expr) Stmt{
  return &Print{
expression,
  }
}

func (PRINT *Print) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitPrintStmt(PRINT)
}


func (rec *Print) IsType(v interface{}) bool {
  switch v.(type) {
   case *Print: return true
  }
  return false
}
type Return struct {
Keyword *Tok
Value Expr
}

func NewReturn(keyword *Tok, value Expr) Stmt{
  return &Return{
keyword,value,
  }
}

func (RETURN *Return) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitReturnStmt(RETURN)
}


func (rec *Return) IsType(v interface{}) bool {
  switch v.(type) {
   case *Return: return true
  }
  return false
}
type Var struct {
Name *Tok
Initializer Expr
}

func NewVar(name *Tok, initializer Expr) Stmt{
  return &Var{
name,initializer,
  }
}

func (VAR *Var) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitVarStmt(VAR)
}


func (rec *Var) IsType(v interface{}) bool {
  switch v.(type) {
   case *Var: return true
  }
  return false
}
type While struct {
Condition Expr
Body Stmt
}

func NewWhile(condition Expr, body Stmt) Stmt{
  return &While{
condition,body,
  }
}

func (WHILE *While) Accept(visitor VisitorStmt) (interface{}, error) {
  return visitor.visitWhileStmt(WHILE)
}


func (rec *While) IsType(v interface{}) bool {
  switch v.(type) {
   case *While: return true
  }
  return false
}
