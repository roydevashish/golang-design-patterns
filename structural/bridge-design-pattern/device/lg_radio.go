package device

import "fmt"

type LGRadio struct{}

func NewLGRadio() *LGRadio {
	return &LGRadio{}
}

func (l *LGRadio) TurnOff() {
	fmt.Println("turning off lg radio device")
}

func (l *LGRadio) TurnOn() {
	fmt.Println("turning on lg radio device")
}

func (l *LGRadio) SetChannel(channel int) {
	fmt.Println("set channel of lg radio device to", channel)
}
