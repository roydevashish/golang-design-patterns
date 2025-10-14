package device

import "fmt"

type SonyRadio struct{}

func NewSonyRadio() *SonyRadio {
	return &SonyRadio{}
}

func (s *SonyRadio) TurnOff() {
	fmt.Println("turning off sony radio device")
}

func (s *SonyRadio) TurnOn() {
	fmt.Println("turning on sony radio device")
}

func (s *SonyRadio) SetChannel(channel int) {
	fmt.Println("set channel of sony radio device to", channel)
}
