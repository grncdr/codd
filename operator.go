package codd

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

func Comparator(op string) func(Expression, Expression) Comparison {
	return func(left Expression, right Expression) Comparison {
		return Comparison{BinaryExpr{op, left, right}}
	}
}

func BooleanOp(op string) func(Boolean, Boolean) Boolean {
	return func(left Boolean, right Boolean) Boolean {
		return Comparison{BinaryExpr{op, left, right}}
	}
}

func Arith(op string) func(Expression, Expression) Arithmetic {
	return func(left Expression, right Expression) Arithmetic {
		return Arithmetic{BinaryExpr{op, left, right}}
	}
}

func TrinaryOp(op, middle string) func(Expression, Expression, Expression) TrinaryExpr {
	return func(a Expression, b Expression, c Expression) TrinaryExpr {
		return TrinaryExpr{op, middle, a, b, c}
	}
}
