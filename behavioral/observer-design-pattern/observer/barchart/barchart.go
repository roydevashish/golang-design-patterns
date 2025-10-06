package barchart

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/observer/subject"
)

type Barchart struct {
	ds *subject.DataSource
}

func NewBarchart(ds *subject.DataSource) *Barchart {
	return &Barchart{
		ds: ds,
	}
}

func (b *Barchart) Update() {
	fmt.Println("render barchart for", b.ds.GetValues())
}
