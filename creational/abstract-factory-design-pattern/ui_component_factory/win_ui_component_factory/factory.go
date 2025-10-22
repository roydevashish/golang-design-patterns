package winuicomponentfactory

import (
	"github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component/button"
	"github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component/checkbox"
	uiwin "github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component/ui_win"
)

type WinUIComponentFactory struct{}

func NewWinUIComponentFactory() *WinUIComponentFactory {
	return &WinUIComponentFactory{}
}

func (m *WinUIComponentFactory) CreateButton() button.IButton {
	return uiwin.NewWinButton()
}

func (m *WinUIComponentFactory) CreateCheckbox() checkbox.ICheckbox {
	return uiwin.NewWinCheckbox()
}
