package expression

import "github.com/roydevashish/golang-design-patterns/interpreter/context"

type Expression interface {
	Interpret(context context.Context) int
}
