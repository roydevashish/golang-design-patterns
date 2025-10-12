package microphone

type Microphone struct {
	price float32
}

func NewMicrophone(p float32) *Microphone {
	return &Microphone{
		price: p,
	}
}

func (m *Microphone) GetPrice() float32 {
	return m.price
}
