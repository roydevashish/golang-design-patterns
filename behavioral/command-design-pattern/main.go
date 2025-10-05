package main

import (
	"github.com/roydevashish/golang-design-patterns/command/command"
	"github.com/roydevashish/golang-design-patterns/command/light"
	remotecontrol "github.com/roydevashish/golang-design-patterns/command/remote_control"
)

func main() {
	l := light.NewLight()
	r := remotecontrol.NewRemoteControl(command.NewTurnOnCommand(l))
	r.PressButton()
	r.SetCommand(command.NewDimCommand(l))
	r.PressButton()
	r.SetCommand(command.NewTurnOffCommand(l))
	r.PressButton()
}
