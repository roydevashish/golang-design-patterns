package main

import (
	ostype "github.com/roydevashish/golang-design-patterns/abstract-factory/os_type"
	uicomponentfactory "github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component_factory"
	macuicomponentfactory "github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component_factory/mac_ui_component_factory"
	winuicomponentfactory "github.com/roydevashish/golang-design-patterns/abstract-factory/ui_component_factory/win_ui_component_factory"
	usersettingsform "github.com/roydevashish/golang-design-patterns/abstract-factory/user_settings_form"
)

func main() {
	os := ostype.MAC
	var uicomponentfactory uicomponentfactory.IUIComponentFactory

	if os == ostype.MAC {
		uicomponentfactory = macuicomponentfactory.NewMacUIComponentFactory()
	} else if os == ostype.WIN {
		uicomponentfactory = winuicomponentfactory.NewWinUIComponentFactory()
	} else {
		panic("unsupported os")
	}

	usersettingsform.NewUserSettingsForm().Render(uicomponentfactory)
}
