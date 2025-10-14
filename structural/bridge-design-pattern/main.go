package main

import (
	"github.com/roydevashish/golang-design-patterns/bridge/device"
	remotecontrol "github.com/roydevashish/golang-design-patterns/bridge/remote_control"
)

func main() {
	lgRemoteControl := remotecontrol.NewRemoteControl(device.NewLGRadio())
	lgRemoteControl.TurnOn()
	lgRemoteControl.TurnOff()

	sonyAdvRemoteControl := remotecontrol.NewAdvanceRemoteControl(device.NewSonyRadio())
	sonyAdvRemoteControl.TurnOn()
	sonyAdvRemoteControl.SetChannel(10)
	sonyAdvRemoteControl.TurnOff()
}
