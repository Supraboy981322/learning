package main

type Expr interface {
  Accept(VisitorExpr) (interface{}, error)
  IsType(interface{}) bool
}

type VisitorExpr interface {
  visitAssignExpr(*Assign) (interface{}, error)
  visitBinaryExpr(*Binary) (interface{}, error)
  visitCallExpr(*Call) (interface{}, error)
  visitGetExpr(*Get) (interface{}, error)
  visitGroupingExpr(*Grouping) (interface{}, error)
  visitLiteralExpr(*Literal) (interface{}, error)
  visitLogicalExpr(*Logical) (interface{}, error)
  visitSetExpr(*Set) (interface{}, error)
  visitSuperExpr(*Super) (interface{}, error)
  visitThisExpr(*This) (interface{}, error)
  visitUnaryExpr(*Unary) (interface{}, error)
  visitVariableExpr(*Variable) (interface{}, error)
}

type Assign struct {
Name *Tok
Value Expr
}

func NewAssign(name *Tok, value Expr) Expr{
  return &Assign{
name,value,
  }
}

func (ASSIGN *Assign) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitAssignExpr(ASSIGN)
}


func (rec *Assign) IsType(v interface{}) bool {
  switch v.(type) {
   case *Assign: return true
  }
  return false
}
type Binary struct {
Left Expr
Operator *Tok
Right Expr
}

func NewBinary(left Expr, operator *Tok, right Expr) Expr{
  return &Binary{
left,operator,right,
  }
}

func (BINARY *Binary) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitBinaryExpr(BINARY)
}


func (rec *Binary) IsType(v interface{}) bool {
  switch v.(type) {
   case *Binary: return true
  }
  return false
}
type Call struct {
Callee Expr
Paren *Tok
Arguments []Expr
}

func NewCall(callee Expr, paren *Tok, arguments []Expr) Expr{
  return &Call{
callee,paren,arguments,
  }
}

func (CALL *Call) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitCallExpr(CALL)
}


func (rec *Call) IsType(v interface{}) bool {
  switch v.(type) {
   case *Call: return true
  }
  return false
}
type Get struct {
Object Expr
Name *Tok
}

func NewGet(object Expr, name *Tok) Expr{
  return &Get{
object,name,
  }
}

func (GET *Get) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitGetExpr(GET)
}


func (rec *Get) IsType(v interface{}) bool {
  switch v.(type) {
   case *Get: return true
  }
  return false
}
type Grouping struct {
Expression Expr
}

func NewGrouping(expression Expr) Expr{
  return &Grouping{
expression,
  }
}

func (GROUPING *Grouping) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitGroupingExpr(GROUPING)
}


func (rec *Grouping) IsType(v interface{}) bool {
  switch v.(type) {
   case *Grouping: return true
  }
  return false
}
type Literal struct {
Value interface{}
}

func NewLiteral(value interface{}) Expr{
  return &Literal{
value,
  }
}

func (LITERAL *Literal) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitLiteralExpr(LITERAL)
}


func (rec *Literal) IsType(v interface{}) bool {
  switch v.(type) {
   case *Literal: return true
  }
  return false
}
type Logical struct {
Left Expr
Operator *Tok
Right Expr
}

func NewLogical(left Expr, operator *Tok, right Expr) Expr{
  return &Logical{
left,operator,right,
  }
}

func (LOGICAL *Logical) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitLogicalExpr(LOGICAL)
}


func (rec *Logical) IsType(v interface{}) bool {
  switch v.(type) {
   case *Logical: return true
  }
  return false
}
type Set struct {
Object Expr
Name *Tok
Value Expr
}

func NewSet(object Expr, name *Tok, value Expr) Expr{
  return &Set{
object,name,value,
  }
}

func (SET *Set) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitSetExpr(SET)
}


func (rec *Set) IsType(v interface{}) bool {
  switch v.(type) {
   case *Set: return true
  }
  return false
}
type Super struct {
Keyword *Tok
Method *Tok
}

func NewSuper(keyword *Tok, method *Tok) Expr{
  return &Super{
keyword,method,
  }
}

func (SUPER *Super) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitSuperExpr(SUPER)
}


func (rec *Super) IsType(v interface{}) bool {
  switch v.(type) {
   case *Super: return true
  }
  return false
}
type This struct {
Keyword *Tok
}

func NewThis(keyword *Tok) Expr{
  return &This{
keyword,
  }
}

func (THIS *This) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitThisExpr(THIS)
}


func (rec *This) IsType(v interface{}) bool {
  switch v.(type) {
   case *This: return true
  }
  return false
}
type Unary struct {
Operator *Tok
Right Expr
}

func NewUnary(operator *Tok, right Expr) Expr{
  return &Unary{
operator,right,
  }
}

func (UNARY *Unary) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitUnaryExpr(UNARY)
}


func (rec *Unary) IsType(v interface{}) bool {
  switch v.(type) {
   case *Unary: return true
  }
  return false
}
type Variable struct {
Name *Tok
}

func NewVariable(name *Tok) Expr{
  return &Variable{
name,
  }
}

func (VARIABLE *Variable) Accept(visitor VisitorExpr) (interface{}, error) {
  return visitor.visitVariableExpr(VARIABLE)
}


func (rec *Variable) IsType(v interface{}) bool {
  switch v.(type) {
   case *Variable: return true
  }
  return false
}
