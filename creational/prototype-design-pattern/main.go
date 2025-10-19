package main

import "github.com/roydevashish/golang-design-patterns/prototype/shape"

func main() {
	circle := shape.NewCircle()
	circle.Draw()
	circle.Radius = 12

	rectangle := shape.NewRectangle()
	rectangle.Draw()
	rectangle.Height = 12
	rectangle.Width = 20

	shapeAction := shape.NewShapeActions()
	newCircle := shapeAction.Duplicate(circle)
	newCircle.Draw()

	newRectangle := shapeAction.Duplicate(rectangle)
	newRectangle.Draw()
}
