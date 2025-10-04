package shoppinglist

import "github.com/roydevashish/golang-design-patterns/iterator/iterator"

type ShoppingList struct {
	list []string
}

func NewShoppingList() *ShoppingList {
	return &ShoppingList{
		list: make([]string, 0),
	}
}

func (s *ShoppingList) Push(item string) {
	s.list = append(s.list, item)
}

func (s *ShoppingList) Pop() {
	s.list = s.list[:len(s.list)-1]
}

func (s *ShoppingList) CreateIterator() iterator.Iterator {
	return NewShoppingListIterator(s)
}
