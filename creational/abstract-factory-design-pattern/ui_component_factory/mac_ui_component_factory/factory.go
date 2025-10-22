package macuicomponentfactory

import (
	"github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component/button"
	"github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component/checkbox"
	uimac "github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component/ui_mac"
)

type MacUIComponentFactory struct{}

func NewMacUIComponentFactory() *MacUIComponentFactory {
	return &MacUIComponentFactory{}
}

func (m *MacUIComponentFactory) CreateButton() button.IButton {
	return uimac.NewMacButton()
}

func (m *MacUIComponentFactory) CreateCheckbox() checkbox.ICheckbox {
	return uimac.NewMacCheckbox()
}
