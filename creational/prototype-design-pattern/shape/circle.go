package shape

import "fmt"

type Circle struct {
	Radius int
}

func NewCircle() *Circle {
	return &Circle{
		Radius: 5,
	}
}

func (c *Circle) Draw() {
	fmt.Println("drawing circle")
}

func (c *Circle) Duplicate() Shape {
	newCircle := NewCircle()
	newCircle.Radius = c.Radius
	return newCircle
}
