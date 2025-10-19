package shape

import "fmt"

type ShapeActions struct{}

func NewShapeActions() *ShapeActions {
	return &ShapeActions{}
}

func (s *ShapeActions) Duplicate(shape Shape) Shape {
	fmt.Println("duplicating shape")
	return shape.Duplicate()
}
