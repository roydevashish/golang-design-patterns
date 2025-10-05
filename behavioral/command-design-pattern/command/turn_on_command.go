package command

import "github.com/roydevashish/golang-design-patterns/command/light"

type TurnOnCommand struct {
	light *light.Light
}

func NewTurnOnCommand(light *light.Light) *TurnOnCommand {
	return &TurnOnCommand{
		light: light,
	}
}

func (t *TurnOnCommand) Execute() {
	t.light.TurnOn()
}
