package sheet

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/observer/subject"
)

type Sheet struct {
	ds *subject.DataSource
}

func NewSheet(ds *subject.DataSource) *Sheet {
	return &Sheet{
		ds: ds,
	}
}

func (s *Sheet) calculateSum(values []int) int {
	sum := 0

	for _, v := range values {
		sum += v
	}

	return sum
}

func (s *Sheet) Update() {
	fmt.Println("calculate sheet sum:", s.calculateSum(s.ds.GetValues()))
}
