package uimac

import (
	"fmt"
)

type MacButton struct{}

func NewMacButton() *MacButton {
	return &MacButton{}
}

func (m *MacButton) Render() {
	fmt.Println("mac: button")
}

type MacCheckbox struct{}

func NewMacCheckbox() *MacCheckbox {
	return &MacCheckbox{}
}

func (m *MacCheckbox) Render() {
	fmt.Println("mac: checkbox")
}
