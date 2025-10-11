package expression

import "github.com/roydevashish/golang-design-patterns/interpreter/context"

type SubtractionExpression struct {
	left  Expression
	right Expression
}

func NewSubtractionExpression(left, right Expression) *SubtractionExpression {
	return &SubtractionExpression{
		left:  left,
		right: right,
	}
}

func (s *SubtractionExpression) Interpret(context context.Context) int {
	return s.left.Interpret(context) - s.right.Interpret(context)
}
