package command

import "github.com/roydevashish/golang-design-patterns/command/light"

type TurnOffCommand struct {
	light *light.Light
}

func NewTurnOffCommand(light *light.Light) *TurnOffCommand {
	return &TurnOffCommand{
		light: light,
	}
}

func (t *TurnOffCommand) Execute() {
	t.light.TurnOff()
}
