package inventory

import "fmt"

type Inventory struct{}

func NewInventory() *Inventory {
	return &Inventory{}
}

func (i *Inventory) CheckInventory(itemId string) bool {
	return true
}

func (i *Inventory) ReduceInventory(itemId string, amount int) {
	fmt.Println("reducing inventory of", itemId, "by", amount)
}
