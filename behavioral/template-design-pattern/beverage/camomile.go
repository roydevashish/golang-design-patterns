package beverage

import (
	"fmt"
)

type Camomile struct{}

func NewCamomile() *Camomile {
	return &Camomile{}
}

func (c *Camomile) Prepare() {
	c.brew()
}

func (c *Camomile) brew() {
	fmt.Println("brewing camomile for 3 min")
}
