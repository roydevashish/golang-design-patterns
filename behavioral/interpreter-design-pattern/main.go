package main

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/interpreter/context"
	"github.com/roydevashish/golang-design-patterns/interpreter/interpreter"
)

func main() {
	var input string = "1 + 2 * 3"
	context := context.NewContext()
	interpreter := interpreter.NewInterpreter(*context)
	result := interpreter.Interpret(input)
	fmt.Println(result)
}
