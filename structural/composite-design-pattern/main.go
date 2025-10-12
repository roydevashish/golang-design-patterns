package main

import (
	"fmt"

	"github.com/roydevashish/golang-design-patterns/composite/item/box"
	"github.com/roydevashish/golang-design-patterns/composite/item/keyboard"
	"github.com/roydevashish/golang-design-patterns/composite/item/microphone"
	"github.com/roydevashish/golang-design-patterns/composite/item/mouse"
)

func main() {
	box0 := box.NewBox()

	box1 := box.NewBox()
	box1.AddItem(microphone.NewMicrophone(4000.00))

	box2 := box.NewBox()

	box3 := box.NewBox()
	box3.AddItem(mouse.NewMouse(100.40))

	box4 := box.NewBox()
	box4.AddItem(keyboard.NewKeyBoard(500.33))

	box2.AddItem(box3)
	box2.AddItem(box4)

	fmt.Println(box2.GetPrice())

	box0.AddItem(box1)
	box0.AddItem(box2)

	fmt.Println(box0.GetPrice())
}
