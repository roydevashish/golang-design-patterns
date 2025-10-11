package interpreter

import (
	"github.com/roydevashish/golang-design-patterns/interpreter/context"
	"github.com/roydevashish/golang-design-patterns/interpreter/expression"
)

type Interpreter struct {
	context context.Context
}

func NewInterpreter(context context.Context) *Interpreter {
	return &Interpreter{
		context: context,
	}
}

func (i *Interpreter) Interpret(exp string) int {
	var expTree expression.Expression = i.BuildExpressionTree(exp)
	return expTree.Interpret(i.context)
}

func (i *Interpreter) BuildExpressionTree(input string) expression.Expression {
	input = "1 + 2 * 3"
	expressionTree := expression.NewAdditionExpression(
		expression.NewNumberExpression(1),
		expression.NewMultiplicationExpression(expression.NewNumberExpression(2), expression.NewNumberExpression(3)),
	)

	return expressionTree
}
