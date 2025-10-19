package shape

import "fmt"

type Rectangle struct {
	Height int
	Width  int
}

func NewRectangle() *Rectangle {
	return &Rectangle{
		Height: 5,
		Width:  10,
	}
}

func (r *Rectangle) Draw() {
	fmt.Println("drawing rectangle")
}

func (r *Rectangle) Duplicate() Shape {
	newRectangle := NewRectangle()
	newRectangle.Height = r.Height
	newRectangle.Width = r.Width
	return newRectangle
}
