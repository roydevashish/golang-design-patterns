package uiwin

import (
	"fmt"
)

type WinButton struct{}

func NewWinButton() *WinButton {
	return &WinButton{}
}

func (w *WinButton) Render() {
	fmt.Println("win: button")
}

type WinCheckbox struct{}

func NewWinCheckbox() *WinCheckbox {
	return &WinCheckbox{}
}

func (w *WinCheckbox) Render() {
	fmt.Println("win: checkbox")
}
