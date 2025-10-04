package shoppinglist

type ListIterator struct {
	collection *ShoppingList
	idx        int
}

func NewShoppingListIterator(col *ShoppingList) *ListIterator {
	return &ListIterator{
		collection: col,
		idx:        0,
	}
}

func (s *ListIterator) Current() string {
	return s.collection.list[s.idx]
}

func (s *ListIterator) Next() {
	if s.HasNext() {
		s.idx++
	}
}

func (s *ListIterator) HasNext() bool {
	return s.idx < len(s.collection.list)
}
