package codd

import (
//"fmt"
)

var (
	EQ  = Comparator(" = ")
	NE  = Comparator(" != ")
	GT  = Comparator(" > ")
	GTE = Comparator(" >= ")
	LT  = Comparator(" < ")
	LTE = Comparator(" <= ")
	Add = Arith(" + ")
	Mul = Arith(" * ")
	Sub = Arith(" - ")
	Div = Arith(" / ")
	And = BooleanOp(" AND ")
	Or  = BooleanOp(" OR ")

	Between = TrinaryOp(" BETWEEN ", " AND ")
)

var OperatorPrecedence = map[string]int{
	" * ":  5,
	" / ":  5,
	" + ":  10,
	" - ":  10,
	" != ": 15,
	" = ":  15,
	" > ":  15,
	" < ":  15,
	" >= ": 15,
	" <= ": 15,
}

// Operand is a marker interface that indicates a Node can be used in comparisons,
// arithmetic, and other expressions. It doesn't make any guarantees that such
// comparisons will be sensible.
type Expression interface {
	Node
	ReferencesColumns
	Precedence() int
}

type BinaryExpr struct {
	op    string
	left  Expression
	right Expression
}

func (expr BinaryExpr) Kind() string {
	return "BinaryExpr"
}

// Shared implementation of compilation for all binary expressions
func (expr BinaryExpr) SQL(compiler Compiler) {
	expr.pushOperand(compiler, expr.left)
	compiler.PushText(expr.op)
	expr.pushOperand(compiler, expr.right)
}

func (expr BinaryExpr) pushOperand(compiler Compiler, child Expression) {
	needsParens := expr.Precedence() < child.Precedence()
	if needsParens {
		compiler.PushText("(")
	}
	compiler.Push(child)
	if needsParens {
		compiler.PushText(")")
	}
}

func (expr BinaryExpr) Precedence() int {
	return OperatorPrecedence[expr.op]
}

func (expr BinaryExpr) ReferencedColumns() []Column {
	result := expr.left.ReferencedColumns()
	result = append(result, expr.right.ReferencedColumns()...)
	return result
}

func (expr BinaryExpr) As(name Name) ExprAlias {
	return ExprAlias{expr, name}
}

type Comparison struct {
	BinaryExpr
}

type Arithmetic struct {
	BinaryExpr
}

func (a Arithmetic) Kind() string {
	return "Arithmetic"
}

func (expr Arithmetic) As(name Name) ExprAlias {
	return ExprAlias{expr, name}
}

func (a Arithmetic) IsNumeric() {}

// Boolean is a marker interface indicating an operand has a boolean value.
type Boolean interface {
	Expression
	IsBoolean()
}

// Numeric is a marker interface indicating an operand can be treated as a number.
type Numeric interface {
	Expression
	IsNumeric()
}

func Comparator(op string) func(Expression, Expression) Comparison {
	return func(left Expression, right Expression) Comparison {
		return Comparison{BinaryExpr{op, left, right}}
	}
}

func Arith(op string) func(Expression, Expression) Arithmetic {
	return func(left Expression, right Expression) Arithmetic {
		return Arithmetic{BinaryExpr{op, left, right}}
	}
}

func BooleanOp(op string) func(Boolean, Boolean) Boolean {
	return func(left Boolean, right Boolean) Boolean {
		return Comparison{BinaryExpr{op, left, right}}
	}
}

func (cmp Comparison) IsBoolean() {}

func (cmp Comparison) Kind() string {
	return "Comparison"
}

type TrinaryExpr struct {
	op     string
	middle string
	a      Expression
	b      Expression
	c      Expression
}

func TrinaryOp(op, middle string) func(Expression, Expression, Expression) TrinaryExpr {
	return func(a Expression, b Expression, c Expression) TrinaryExpr {
		return TrinaryExpr{op, middle, a, b, c}
	}
}

func (cmp TrinaryExpr) SQL(builder Compiler) {
	builder.Push(cmp.a)
	builder.PushText(cmp.op)
	builder.Push(cmp.b)
	if cmp.middle != "" {
		builder.PushText(cmp.middle)
	}
	builder.Push(cmp.c)
}

type ExprAlias struct {
	expr Expression
	name Name
}

func (a ExprAlias) Kind() string {
	return "ExprAlias"
}

func (a ExprAlias) SQL(compiler Compiler) {
	compiler.Push(a.expr)
	compiler.PushText(" AS ")
	compiler.Push(a.name)
}

func (a ExprAlias) DBName() Name {
	return a.name
}

func (a ExprAlias) ReferencedColumns() []Column {
	return a.expr.ReferencedColumns()
}
