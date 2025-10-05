package command

import "github.com/roydevashish/golang-design-patterns/command/light"

type DimCommand struct {
	light *light.Light
}

func NewDimCommand(light *light.Light) *DimCommand {
	return &DimCommand{
		light: light,
	}
}

func (d *DimCommand) Execute() {
	d.light.Dim()
}
