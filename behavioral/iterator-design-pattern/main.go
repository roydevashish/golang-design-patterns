package main

import (
	"fmt"

	shoppinglist "github.com/roydevashish/golang-design-patterns/iterator/shopping_list"
)

func main() {
	s := shoppinglist.NewShoppingList()
	s.Push("Milk")
	s.Push("Bread")
	s.Push("Cream")

	it := s.CreateIterator()
	for it.HasNext() {
		fmt.Println(it.Current())
		it.Next()
	}

	s.Push("Jam")
	s.Push("Butter")

	// s.Pop()

	for it.HasNext() {
		fmt.Println(it.Current())
		it.Next()
	}
}
