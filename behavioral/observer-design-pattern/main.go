package main

import (
	"github.com/roydevashish/golang-design-patterns/observer/observer/barchart"
	"github.com/roydevashish/golang-design-patterns/observer/observer/sheet"
	"github.com/roydevashish/golang-design-patterns/observer/subject"
)

func main() {
	ds := subject.NewDataSource()

	s := sheet.NewSheet(ds)
	b := barchart.NewBarchart(ds)

	ds.AddObserver(s)
	ds.AddObserver(b)

	ds.SetValues([]int{1, 2, 3, 4, 5})
	ds.SetValues([]int{1, 4, 3, 1, 0})
	ds.SetValues([]int{1, 1, 3, 1, 0})
}
