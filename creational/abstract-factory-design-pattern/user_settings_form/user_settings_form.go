package usersettingsform

import uicomponentfactory "github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component_factory"

type UserSettingsForm struct{}

func NewUserSettingsForm() *UserSettingsForm {
	return &UserSettingsForm{}
}

func (u *UserSettingsForm) Render(uicomponentfactory uicomponentfactory.IUIComponentFactory) {
	uicomponentfactory.CreateButton().Render()
	uicomponentfactory.CreateCheckbox().Render()
}
