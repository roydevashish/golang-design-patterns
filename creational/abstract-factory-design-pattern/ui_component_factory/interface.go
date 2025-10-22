package uicomponentfactory

import (
	"github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component/button"
	"github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component/checkbox"
)

type IUIComponentFactory interface {
	CreateButton() button.IButton
	CreateCheckbox() checkbox.ICheckbox
}
