package expression

import "github.com/roydevashish/golang-design-patterns/interpreter/context"

type MultiplicationExpression struct {
	left  Expression
	right Expression
}

func NewMultiplicationExpression(left, right Expression) *MultiplicationExpression {
	return &MultiplicationExpression{
		left:  left,
		right: right,
	}
}

func (m *MultiplicationExpression) Interpret(context context.Context) int {
	return m.left.Interpret(context) * m.right.Interpret(context)
}
