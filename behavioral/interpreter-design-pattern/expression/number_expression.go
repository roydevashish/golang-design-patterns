package expression

import "github.com/roydevashish/golang-design-patterns/interpreter/context"

type NumberExpression struct {
	num int
}

func NewNumberExpression(num int) *NumberExpression {
	return &NumberExpression{
		num: num,
	}
}

func (n *NumberExpression) Interpret(context context.Context) int {
	return n.num
}
