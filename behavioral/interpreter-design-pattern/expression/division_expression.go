package expression

import "github.com/roydevashish/golang-design-patterns/interpreter/context"

type DivisionExpression struct {
	left  Expression
	right Expression
}

func NewDivisionExpression(left, right Expression) *DivisionExpression {
	return &DivisionExpression{
		left:  left,
		right: right,
	}
}

func (d *DivisionExpression) Interpret(context context.Context) int {
	return d.left.Interpret(context) / d.right.Interpret(context)
}
