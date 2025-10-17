package orderfulfillment

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/facade/inventory"
)

type OrderFulfillment struct {
	inventory *inventory.Inventory
}

func NewOrderFulfillment(inventory *inventory.Inventory) *OrderFulfillment {
	return &OrderFulfillment{
		inventory: inventory,
	}
}

func (o *OrderFulfillment) FulFill(name, address string, items []string) {
	fmt.Println("inserting order into database")
	for _, item := range items {
		o.inventory.ReduceInventory(item, 1)
	}
}
