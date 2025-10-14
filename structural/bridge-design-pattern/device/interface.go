package device

type Device interface {
	TurnOn()
	TurnOff()
	SetChannel(channel int)
}
