package light

import "fmt"

type Light struct{}

func NewLight() *Light {
	return &Light{}
}

func (l *Light) TurnOn() {
	fmt.Println("turn on the light")
}

func (l *Light) TurnOff() {
	fmt.Println("turn off the light")
}

func (l *Light) Dim() {
	fmt.Println("dim the light")
}
