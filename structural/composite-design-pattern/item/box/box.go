package box

import "github.com/roydevashish/golang-design-patterns/composite/item"

type Box struct {
	itemList []item.Item
}

func NewBox() *Box {
	return &Box{}
}

func (b *Box) AddItem(i item.Item) {
	b.itemList = append(b.itemList, i)
}

func (b *Box) GetPrice() float32 {
	var total float32 = 0

	for _, i := range b.itemList {
		total += i.GetPrice()
	}

	return total
}
