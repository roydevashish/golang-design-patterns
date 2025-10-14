package remotecontrol

import "github.com/roydevashish/golang-design-patterns/bridge/device"

type AdvanceRemoteControl struct {
	RemoteControl
}

func NewAdvanceRemoteControl(d device.Device) *AdvanceRemoteControl {
	return &AdvanceRemoteControl{
		*NewRemoteControl(d),
	}
}

func (a *AdvanceRemoteControl) SetChannel(channel int) {
	a.device.SetChannel(channel)
}
