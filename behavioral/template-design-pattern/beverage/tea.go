package beverage

import (
	"fmt"
	"strings"
)

type Tea struct{}

func NewTea() *Tea {
	return &Tea{}
}

func (t *Tea) Prepare() {
	t.brew()
	t.addCondiments()
}

func (t *Tea) brew() {
	fmt.Println("brewing tea for 3 min")
}

func (t *Tea) addCondiments() {
	if t.customerWantsCondiments() {
		fmt.Println("adding lemon to tea")
	}
}

func (t *Tea) customerWantsCondiments() bool {
	fmt.Println("Do you want lemon in your tea y/n:")
	var x string
	fmt.Scanln(&x)
	return strings.ToLower(x) == "y"
}
