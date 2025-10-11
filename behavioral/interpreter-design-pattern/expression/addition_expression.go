package expression

import "github.com/roydevashish/golang-design-patterns/interpreter/context"

type AdditionExpression struct {
	left  Expression
	right Expression
}

func NewAdditionExpression(left, right Expression) *AdditionExpression {
	return &AdditionExpression{
		left:  left,
		right: right,
	}
}

func (a *AdditionExpression) Interpret(context context.Context) int {
	return a.left.Interpret(context) + a.right.Interpret(context)
}
