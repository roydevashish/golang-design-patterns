package remotecontrol

import "github.com/roydevashish/golang-design-patterns/bridge/device"

type RemoteControl struct {
	device device.Device
}

func NewRemoteControl(d device.Device) *RemoteControl {
	return &RemoteControl{
		device: d,
	}
}

func (r *RemoteControl) TurnOn() {
	r.device.TurnOn()
}

func (r *RemoteControl) TurnOff() {
	r.device.TurnOff()
}
