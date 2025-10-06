package beveragemaker

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/template/beverage"
)

type BeverageMaker struct {
	beverage beverage.Beverage
}

func NewBeverageMaker(beverage beverage.Beverage) *BeverageMaker {
	return &BeverageMaker{
		beverage: beverage,
	}
}

func (b *BeverageMaker) SetBeverage(beverage beverage.Beverage) {
	b.beverage = beverage
}

func (b *BeverageMaker) boilWater() {
	fmt.Println("boiling water")
}

func (b *BeverageMaker) pourIntoCup() {
	fmt.Println("pouring boiled water into cup")
}

func (b *BeverageMaker) PrepareBeverage() {
	b.boilWater()
	b.pourIntoCup()
	b.beverage.Prepare()
	fmt.Println()
}
