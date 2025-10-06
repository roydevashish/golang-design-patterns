package beverage

import (
	"fmt"
	"strings"
)

type Coffee struct{}

func NewCoffee() *Coffee {
	return &Coffee{}
}

func (c *Coffee) Prepare() {
	c.brew()
	c.addCondiments()
}

func (c *Coffee) brew() {
	fmt.Println("brewing coffee for 5 min")
}

func (c *Coffee) addCondiments() {
	if c.customerWantsCondiments() {
		fmt.Println("adding cream to coffee")
	}
}

func (c *Coffee) customerWantsCondiments() bool {
	fmt.Println("Do you want cream in your coffee y/n:")
	var x string
	fmt.Scanln(&x)
	return strings.ToLower(x) == "y"
}
