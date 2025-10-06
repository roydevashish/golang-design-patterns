package main

import (
	"github.com/roydevashish/golang-design-patterns/template/beverage"
	beveragemaker "github.com/roydevashish/golang-design-patterns/template/beverage_maker"
)

func main() {
	b := beverage.NewTea()
	bm := beveragemaker.NewBeverageMaker(b)
	bm.PrepareBeverage()
	bm.SetBeverage(beverage.NewCoffee())
	bm.PrepareBeverage()
	bm.SetBeverage(beverage.NewCamomile())
	bm.PrepareBeverage()
}
